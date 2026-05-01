// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package imports

import (
	"database/sql"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"homefeed/internal/workspace"
)

// Service copies local folders into a workspace feed and records the batch in
// hidden metadata.
type Service struct {
	layout workspace.Layout
}

// Result captures the outcome of a folder import.
type Result struct {
	BatchID         int64  `json:"batchId"`
	SourcePath      string `json:"sourcePath"`
	FeedSlug        string `json:"feedSlug"`
	FeedPath        string `json:"feedPath"`
	DestinationPath string `json:"destinationPath"`
}

// NewService creates an import service for a workspace layout.
func NewService(layout workspace.Layout) *Service {
	return &Service{layout: layout}
}

// ImportFolder copies the source directory into the selected feed.
func (s *Service) ImportFolder(sourcePath, feedSlug string) (Result, error) {
	if sourcePath == "" {
		return Result{}, errors.New("source path is required")
	}
	if feedSlug == "" {
		return Result{}, errors.New("feed slug is required")
	}

	if err := ensureWorkspaceDatabaseExists(s.layout.DatabasePath); err != nil {
		return Result{}, err
	}

	sourceAbs, err := filepath.Abs(sourcePath)
	if err != nil {
		return Result{}, fmt.Errorf("resolve source path: %w", err)
	}

	sourceInfo, err := os.Stat(sourceAbs)
	if err != nil {
		return Result{}, fmt.Errorf("stat source folder: %w", err)
	}
	if !sourceInfo.IsDir() {
		return Result{}, fmt.Errorf("source path %s is not a directory", sourceAbs)
	}

	db, err := openDatabase(s.layout.DatabasePath)
	if err != nil {
		return Result{}, err
	}
	defer db.Close()

	feedPath, err := lookupFeedPath(db, feedSlug)
	if err != nil {
		return Result{}, err
	}

	destinationPath := filepath.Join(feedPath, filepath.Base(sourceAbs))
	if _, err := os.Stat(destinationPath); err == nil {
		return Result{}, fmt.Errorf("destination already exists: %s", destinationPath)
	} else if !os.IsNotExist(err) {
		return Result{}, fmt.Errorf("check destination path: %w", err)
	}

	if err := copyTree(sourceAbs, destinationPath); err != nil {
		_ = os.RemoveAll(destinationPath)
		return Result{}, err
	}

	batchID, err := insertImportBatch(db, sourceAbs, feedSlug, feedPath, destinationPath)
	if err != nil {
		_ = os.RemoveAll(destinationPath)
		return Result{}, err
	}

	return Result{
		BatchID:         batchID,
		SourcePath:      sourceAbs,
		FeedSlug:        feedSlug,
		FeedPath:        feedPath,
		DestinationPath: destinationPath,
	}, nil
}

func ensureWorkspaceDatabaseExists(path string) error {
	if path == "" {
		return errors.New("workspace database path is required")
	}

	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("workspace database not found at %s; initialize the workspace first", path)
		}
		return fmt.Errorf("stat workspace database: %w", err)
	}

	return nil
}

func openDatabase(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
		return nil, fmt.Errorf("open sqlite database: %w", err)
	}
	db.SetMaxOpenConns(1)
	db.SetMaxIdleConns(1)

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ping sqlite database: %w", err)
	}

	return db, nil
}

func lookupFeedPath(db *sql.DB, feedSlug string) (string, error) {
	var feedPath string
	err := db.QueryRow(`SELECT path FROM feeds WHERE slug = ?`, feedSlug).Scan(&feedPath)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", fmt.Errorf("feed %q not found", feedSlug)
		}
		return "", fmt.Errorf("lookup feed path: %w", err)
	}

	if _, err := os.Stat(feedPath); err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("feed path %s does not exist", feedPath)
		}
		return "", fmt.Errorf("stat feed path: %w", err)
	}

	return feedPath, nil
}

func insertImportBatch(db *sql.DB, sourcePath, feedSlug, feedPath, destinationPath string) (int64, error) {
	result, err := db.Exec(
		`INSERT INTO import_batches (source_path, feed_slug, feed_path, destination_path) VALUES (?, ?, ?, ?)`,
		sourcePath,
		feedSlug,
		feedPath,
		destinationPath,
	)
	if err != nil {
		return 0, fmt.Errorf("record import batch: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("read import batch id: %w", err)
	}

	return id, nil
}

func copyTree(source, destination string) error {
	sourceInfo, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("stat source tree: %w", err)
	}

	if err := os.MkdirAll(destination, sourceInfo.Mode().Perm()); err != nil {
		return fmt.Errorf("create destination root: %w", err)
	}

	return filepath.WalkDir(source, func(currentPath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}

		relPath, err := filepath.Rel(source, currentPath)
		if err != nil {
			return fmt.Errorf("resolve relative path: %w", err)
		}

		targetPath := destination
		if relPath != "." {
			targetPath = filepath.Join(destination, relPath)
		}

		if entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return fmt.Errorf("stat directory %s: %w", currentPath, err)
			}
			if err := os.MkdirAll(targetPath, info.Mode().Perm()); err != nil {
				return fmt.Errorf("create directory %s: %w", targetPath, err)
			}
			return nil
		}

		info, err := entry.Info()
		if err != nil {
			return fmt.Errorf("stat file %s: %w", currentPath, err)
		}

		if info.Mode()&os.ModeSymlink != 0 {
			linkTarget, err := os.Readlink(currentPath)
			if err != nil {
				return fmt.Errorf("read symlink %s: %w", currentPath, err)
			}
			if err := os.Symlink(linkTarget, targetPath); err != nil {
				return fmt.Errorf("create symlink %s: %w", targetPath, err)
			}
			return nil
		}

		if !info.Mode().IsRegular() {
			return fmt.Errorf("unsupported file type at %s", currentPath)
		}

		if err := copyFile(currentPath, targetPath, info.Mode().Perm()); err != nil {
			return err
		}

		return nil
	})
}

func copyFile(source, destination string, perm fs.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return fmt.Errorf("prepare destination directory: %w", err)
	}

	in, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("open source file %s: %w", source, err)
	}
	defer in.Close()

	out, err := os.OpenFile(destination, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, perm)
	if err != nil {
		return fmt.Errorf("open destination file %s: %w", destination, err)
	}
	defer out.Close()

	if _, err := io.Copy(out, in); err != nil {
		return fmt.Errorf("copy file %s: %w", source, err)
	}

	if err := out.Sync(); err != nil {
		return fmt.Errorf("sync file %s: %w", destination, err)
	}

	return nil
}

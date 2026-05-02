// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package posts

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"homefeed/internal/workspace"
)

// Service creates note posts inside a workspace feed.
type Service struct {
	layout workspace.Layout
}

// Result captures the outcome of a note-post creation.
type Result struct {
	PostID   int64  `json:"postId"`
	Kind     string `json:"kind"`
	FeedSlug string `json:"feedSlug"`
	FeedPath string `json:"feedPath"`
	Title    string `json:"title"`
	Body     string `json:"body"`
	Path     string `json:"path"`
}

// NewService creates a post service for a workspace layout.
func NewService(layout workspace.Layout) *Service {
	return &Service{layout: layout}
}

// CreateNotePost creates a markdown note file in the selected feed and records
// the post in the workspace database.
func (s *Service) CreateNotePost(feedSlug, title, body string) (Result, error) {
	if feedSlug == "" {
		return Result{}, errors.New("feed slug is required")
	}
	if title == "" {
		return Result{}, errors.New("title is required")
	}

	if err := ensureWorkspaceDatabaseExists(s.layout.DatabasePath); err != nil {
		return Result{}, err
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

	filename := noteFilename(title)
	notePath := filepath.Join(feedPath, filename)
	if _, err := os.Stat(notePath); err == nil {
		return Result{}, fmt.Errorf("note already exists: %s", notePath)
	} else if !os.IsNotExist(err) {
		return Result{}, fmt.Errorf("check note path: %w", err)
	}

	if err := writeNoteFile(notePath, title, body); err != nil {
		return Result{}, err
	}

	postID, err := insertNotePost(db, feedSlug, feedPath, title, body, notePath)
	if err != nil {
		_ = os.Remove(notePath)
		return Result{}, err
	}

	return Result{
		PostID:   postID,
		Kind:     "note",
		FeedSlug: feedSlug,
		FeedPath: feedPath,
		Title:    title,
		Body:     body,
		Path:     notePath,
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

func insertNotePost(db *sql.DB, feedSlug, feedPath, title, body, notePath string) (int64, error) {
	result, err := db.Exec(
		`INSERT INTO posts (kind, feed_slug, feed_path, title, body, path) VALUES (?, ?, ?, ?, ?, ?)`,
		"note",
		feedSlug,
		feedPath,
		title,
		body,
		notePath,
	)
	if err != nil {
		return 0, fmt.Errorf("record note post: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("read note post id: %w", err)
	}

	return id, nil
}

func noteFilename(title string) string {
	slug := normalizeSlug(title)
	if slug == "" {
		slug = "note"
	}
	return slug + ".md"
}

func normalizeSlug(text string) string {
	var builder strings.Builder
	lastDash := false

	for _, r := range strings.ToLower(text) {
		switch {
		case unicode.IsLetter(r), unicode.IsDigit(r):
			builder.WriteRune(r)
			lastDash = false
		case unicode.IsSpace(r), r == '-', r == '_':
			if !lastDash && builder.Len() > 0 {
				builder.WriteByte('-')
				lastDash = true
			}
		}
	}

	return strings.Trim(builder.String(), "-")
}

func writeNoteFile(path, title, body string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("prepare note directory: %w", err)
	}

	content := "# " + title + "\n\n"
	if body != "" {
		content += body + "\n"
	}

	if err := os.WriteFile(path, []byte(content), 0o644); err != nil {
		return fmt.Errorf("write note file: %w", err)
	}

	return nil
}

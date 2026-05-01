// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package workspace

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

// FeedDefinition defines one canonical feed.
type FeedDefinition struct {
	Slug string
	Name string
}

// DefaultFeeds are the canonical top-level feeds created on first run.
var DefaultFeeds = []FeedDefinition{
	{Slug: "professional", Name: "Professional"},
	{Slug: "family", Name: "Family"},
	{Slug: "projects", Name: "Projects"},
	{Slug: "personal", Name: "Personal"},
	{Slug: "archive", Name: "Archive"},
}

// Result captures the initialized workspace state.
type Result struct {
	Layout Layout
	Feeds  []FeedDefinition
}

// Initialize creates the workspace, metadata directories, database, and default feeds.
func Initialize(root string) (Result, error) {
	if root == "" {
		return Result{}, errors.New("workspace root is required")
	}

	layout := LayoutForRoot(root)

	if err := ensureDirectories(layout, VisibleFeedPaths(layout.Root)); err != nil {
		return Result{}, err
	}

	db, err := openDatabase(layout.DatabasePath)
	if err != nil {
		return Result{}, err
	}
	defer db.Close()

	if err := applyInitialMigrations(db); err != nil {
		return Result{}, err
	}

	if err := seedDefaultFeeds(db, layout.Root); err != nil {
		return Result{}, err
	}

	return Result{
		Layout: layout,
		Feeds:  append([]FeedDefinition(nil), DefaultFeeds...),
	}, nil
}

func ensureDirectories(layout Layout, visibleFeedPaths []string) error {
	dirs := []string{
		layout.Root,
		layout.MetaRoot,
		layout.MetaDataPath,
		layout.ProfilesPath,
		layout.ImportsPath,
		layout.AgentsPath,
		layout.LogsPath,
		layout.CachePath,
	}
	dirs = append(dirs, visibleFeedPaths...)

	for _, dir := range dirs {
		if err := os.MkdirAll(dir, 0o755); err != nil {
			return fmt.Errorf("create directory %s: %w", dir, err)
		}
	}

	return nil
}

func openDatabase(path string) (*sql.DB, error) {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return nil, fmt.Errorf("prepare database directory: %w", err)
	}

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

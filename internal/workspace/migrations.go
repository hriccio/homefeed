// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package workspace

import (
	"database/sql"
	"fmt"
	"path/filepath"
)

var createWorkspaceStatements = []string{
	`PRAGMA foreign_keys = ON`,
	`CREATE TABLE IF NOT EXISTS schema_migrations (
    version INTEGER PRIMARY KEY,
    applied_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
)`,
	`CREATE TABLE IF NOT EXISTS feeds (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    slug TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    path TEXT NOT NULL UNIQUE,
    created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
)`,
}

func applyInitialMigrations(db *sql.DB) error {
	for _, statement := range createWorkspaceStatements {
		if _, err := db.Exec(statement); err != nil {
			return fmt.Errorf("apply workspace schema: %w", err)
		}
	}

	if _, err := db.Exec(`INSERT OR IGNORE INTO schema_migrations(version) VALUES (1)`); err != nil {
		return fmt.Errorf("record schema migration: %w", err)
	}

	return nil
}

func seedDefaultFeeds(db *sql.DB, root string) error {
	for _, feed := range DefaultFeeds {
		feedPath := filepath.Join(root, feed.Slug)
		if _, err := db.Exec(
			`INSERT OR IGNORE INTO feeds (slug, name, path) VALUES (?, ?, ?)`,
			feed.Slug,
			feed.Name,
			feedPath,
		); err != nil {
			return fmt.Errorf("seed feed %s: %w", feed.Slug, err)
		}
	}

	return nil
}

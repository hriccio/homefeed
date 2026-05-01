// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package workspace

import (
	"database/sql"
	"fmt"
	"path/filepath"

	homefeedmigrations "homefeed/migrations"
)

func applyInitialMigrations(db *sql.DB) error {
	statements, err := homefeedmigrations.WorkspaceStatements()
	if err != nil {
		return fmt.Errorf("load workspace migrations: %w", err)
	}

	for _, statement := range statements {
		if _, err := db.Exec(statement); err != nil {
			return fmt.Errorf("apply workspace schema: %w", err)
		}
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

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server_test

import (
	"database/sql"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"

	"homefeed/internal/workspace"
)

func TestDefaultWorkspaceRootUsesHomeDirectory(t *testing.T) {
	got := workspace.DefaultWorkspaceRoot("/tmp/test-home")
	want := filepath.Join("/tmp/test-home", "Homefeed")
	if got != want {
		t.Fatalf("default workspace root = %q, want %q", got, want)
	}
}

func TestInitializeWorkspaceCreatesWorkspaceAndDatabase(t *testing.T) {
	root := filepath.Join(t.TempDir(), "Homefeed")

	result, err := workspace.Initialize(root)
	if err != nil {
		t.Fatalf("initialize workspace: %v", err)
	}

	assertDirectoryExists(t, result.Layout.Root)
	assertDirectoryExists(t, result.Layout.MetaRoot)
	assertDirectoryExists(t, result.Layout.MetaDataPath)
	assertDirectoryExists(t, result.Layout.ProfilesPath)
	assertDirectoryExists(t, result.Layout.ImportsPath)
	assertDirectoryExists(t, result.Layout.AgentsPath)
	assertDirectoryExists(t, result.Layout.LogsPath)
	assertDirectoryExists(t, result.Layout.CachePath)
	assertFileExists(t, result.Layout.DatabasePath)

	gotFeeds := readFeedSlugs(t, result.Layout.DatabasePath)
	wantFeeds := []string{"archive", "family", "personal", "professional", "projects"}
	if !reflect.DeepEqual(gotFeeds, wantFeeds) {
		t.Fatalf("feeds = %v, want %v", gotFeeds, wantFeeds)
	}

	for _, feed := range workspace.DefaultFeeds {
		feedPath := filepath.Join(root, feed.Slug)
		assertDirectoryExists(t, feedPath)
		entries, err := os.ReadDir(feedPath)
		if err != nil {
			t.Fatalf("read feed directory %s: %v", feedPath, err)
		}
		if len(entries) != 0 {
			t.Fatalf("feed directory %s is not empty: %d entries", feedPath, len(entries))
		}
	}
}

func TestInitializeWorkspaceIsIdempotent(t *testing.T) {
	root := filepath.Join(t.TempDir(), "Homefeed")

	first, err := workspace.Initialize(root)
	if err != nil {
		t.Fatalf("first initialize: %v", err)
	}

	second, err := workspace.Initialize(root)
	if err != nil {
		t.Fatalf("second initialize: %v", err)
	}

	if first.Layout.DatabasePath != second.Layout.DatabasePath {
		t.Fatalf("database path changed between runs: %q vs %q", first.Layout.DatabasePath, second.Layout.DatabasePath)
	}

	gotFeeds := readFeedSlugs(t, second.Layout.DatabasePath)
	wantFeeds := []string{"archive", "family", "personal", "professional", "projects"}
	if !reflect.DeepEqual(gotFeeds, wantFeeds) {
		t.Fatalf("feeds after second run = %v, want %v", gotFeeds, wantFeeds)
	}

	assertVisibleFolderSet(t, root, []string{"archive", "family", "personal", "professional", "projects", ".homefeed"})
}

func TestDatabasePathsMatchVisibleFeeds(t *testing.T) {
	root := filepath.Join(t.TempDir(), "Homefeed")

	result, err := workspace.Initialize(root)
	if err != nil {
		t.Fatalf("initialize workspace: %v", err)
	}

	db := openTestDB(t, result.Layout.DatabasePath)
	defer db.Close()

	rows, err := db.Query(`SELECT slug, path FROM feeds ORDER BY slug`)
	if err != nil {
		t.Fatalf("query feeds: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var slug, path string
		if err := rows.Scan(&slug, &path); err != nil {
			t.Fatalf("scan row: %v", err)
		}
		if path != filepath.Join(root, slug) {
			t.Fatalf("feed %s path = %q, want %q", slug, path, filepath.Join(root, slug))
		}
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows error: %v", err)
	}
}

func openTestDB(t *testing.T, path string) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite", path)
	if err != nil {
		t.Fatalf("open sqlite database: %v", err)
	}
	return db
}

func readFeedSlugs(t *testing.T, dbPath string) []string {
	t.Helper()

	db := openTestDB(t, dbPath)
	defer db.Close()

	rows, err := db.Query(`SELECT slug FROM feeds ORDER BY slug`)
	if err != nil {
		t.Fatalf("query feeds: %v", err)
	}
	defer rows.Close()

	var slugs []string
	for rows.Next() {
		var slug string
		if err := rows.Scan(&slug); err != nil {
			t.Fatalf("scan feed slug: %v", err)
		}
		slugs = append(slugs, slug)
	}
	if err := rows.Err(); err != nil {
		t.Fatalf("rows error: %v", err)
	}

	sort.Strings(slugs)
	return slugs
}

func assertDirectoryExists(t *testing.T, path string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat directory %s: %v", path, err)
	}
	if !info.IsDir() {
		t.Fatalf("%s is not a directory", path)
	}
}

func assertFileExists(t *testing.T, path string) {
	t.Helper()

	info, err := os.Stat(path)
	if err != nil {
		t.Fatalf("stat file %s: %v", path, err)
	}
	if info.IsDir() {
		t.Fatalf("%s is a directory, want file", path)
	}
}

func assertVisibleFolderSet(t *testing.T, root string, want []string) {
	t.Helper()

	entries, err := os.ReadDir(root)
	if err != nil {
		t.Fatalf("read root directory: %v", err)
	}

	got := make([]string, 0, len(entries))
	for _, entry := range entries {
		got = append(got, entry.Name())
	}
	sort.Strings(got)
	sort.Strings(want)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("root entries = %v, want %v", got, want)
	}
}

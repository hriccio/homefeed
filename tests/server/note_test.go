// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server_test

import (
	"os"
	"path/filepath"
	"testing"

	"homefeed/internal/posts"
	"homefeed/internal/workspace"
)

func TestCreateNotePostCreatesFileAndRecord(t *testing.T) {
	root := filepath.Join(t.TempDir(), "Homefeed")

	result, err := workspace.Initialize(root)
	if err != nil {
		t.Fatalf("initialize workspace: %v", err)
	}

	service := posts.NewService(result.Layout)
	noteResult, err := service.CreateNotePost("projects", "My First Note", "hello from homefeed")
	if err != nil {
		t.Fatalf("create note post: %v", err)
	}

	wantPath := filepath.Join(root, "projects", "my-first-note.md")
	if noteResult.Path != wantPath {
		t.Fatalf("note path = %q, want %q", noteResult.Path, wantPath)
	}
	if noteResult.Kind != "note" {
		t.Fatalf("note kind = %q, want note", noteResult.Kind)
	}

	assertFileContents(t, wantPath, "# My First Note\n\nhello from homefeed\n")

	db := openTestDB(t, result.Layout.DatabasePath)
	defer db.Close()

	var count int
	if err := db.QueryRow(`SELECT COUNT(*) FROM posts`).Scan(&count); err != nil {
		t.Fatalf("count posts: %v", err)
	}
	if count != 1 {
		t.Fatalf("post count = %d, want 1", count)
	}

	var kind, feedSlug, feedPath, title, body, path string
	if err := db.QueryRow(
		`SELECT kind, feed_slug, feed_path, title, body, path FROM posts LIMIT 1`,
	).Scan(&kind, &feedSlug, &feedPath, &title, &body, &path); err != nil {
		t.Fatalf("read post: %v", err)
	}

	if kind != "note" {
		t.Fatalf("stored kind = %q, want note", kind)
	}
	if feedSlug != "projects" {
		t.Fatalf("stored feed slug = %q, want projects", feedSlug)
	}
	if feedPath != filepath.Join(root, "projects") {
		t.Fatalf("stored feed path = %q, want %q", feedPath, filepath.Join(root, "projects"))
	}
	if title != "My First Note" {
		t.Fatalf("stored title = %q, want %q", title, "My First Note")
	}
	if body != "hello from homefeed" {
		t.Fatalf("stored body = %q, want %q", body, "hello from homefeed")
	}
	if path != wantPath {
		t.Fatalf("stored path = %q, want %q", path, wantPath)
	}

	assertVisibleFolderSet(t, root, []string{"archive", "family", "personal", "professional", "projects", ".homefeed"})

	if entries, err := os.ReadDir(filepath.Join(root, ".homefeed")); err != nil {
		t.Fatalf("read hidden metadata root: %v", err)
	} else if len(entries) == 0 {
		t.Fatal("hidden metadata root is empty, want database-backed metadata present")
	}
}

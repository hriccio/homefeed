// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package server_test

import (
	"os"
	"path/filepath"
	"testing"

	"homefeed/internal/imports"
	"homefeed/internal/workspace"
)

func TestImportFolderCopiesSourceFolderIntoWorkspace(t *testing.T) {
	root := filepath.Join(t.TempDir(), "Homefeed")

	result, err := workspace.Initialize(root)
	if err != nil {
		t.Fatalf("initialize workspace: %v", err)
	}

	source := filepath.Join(t.TempDir(), "sample-import")
	if err := os.MkdirAll(filepath.Join(source, "nested"), 0o755); err != nil {
		t.Fatalf("create source tree: %v", err)
	}
	if err := os.WriteFile(filepath.Join(source, "notes.txt"), []byte("hello world"), 0o644); err != nil {
		t.Fatalf("write source file: %v", err)
	}
	if err := os.WriteFile(filepath.Join(source, "nested", "detail.md"), []byte("# detail"), 0o644); err != nil {
		t.Fatalf("write nested source file: %v", err)
	}

	service := imports.NewService(result.Layout)
	importResult, err := service.ImportFolder(source, "projects")
	if err != nil {
		t.Fatalf("import folder: %v", err)
	}

	wantDestination := filepath.Join(root, "projects", filepath.Base(source))
	if importResult.DestinationPath != wantDestination {
		t.Fatalf("destination path = %q, want %q", importResult.DestinationPath, wantDestination)
	}

	assertDirectoryExists(t, wantDestination)
	assertFileContents(t, filepath.Join(wantDestination, "notes.txt"), "hello world")
	assertFileContents(t, filepath.Join(wantDestination, "nested", "detail.md"), "# detail")

	db := openTestDB(t, result.Layout.DatabasePath)
	defer db.Close()

	var batchCount int
	if err := db.QueryRow(`SELECT COUNT(*) FROM import_batches`).Scan(&batchCount); err != nil {
		t.Fatalf("count import batches: %v", err)
	}
	if batchCount != 1 {
		t.Fatalf("import batch count = %d, want 1", batchCount)
	}

	var sourcePath, feedSlug, feedPath, destinationPath string
	if err := db.QueryRow(
		`SELECT source_path, feed_slug, feed_path, destination_path FROM import_batches LIMIT 1`,
	).Scan(&sourcePath, &feedSlug, &feedPath, &destinationPath); err != nil {
		t.Fatalf("read import batch: %v", err)
	}

	if sourcePath != filepath.Clean(source) {
		t.Fatalf("stored source path = %q, want %q", sourcePath, filepath.Clean(source))
	}
	if feedSlug != "projects" {
		t.Fatalf("stored feed slug = %q, want %q", feedSlug, "projects")
	}
	if feedPath != filepath.Join(root, "projects") {
		t.Fatalf("stored feed path = %q, want %q", feedPath, filepath.Join(root, "projects"))
	}
	if destinationPath != wantDestination {
		t.Fatalf("stored destination path = %q, want %q", destinationPath, wantDestination)
	}

	assertVisibleFolderSet(t, root, []string{"archive", "family", "personal", "professional", "projects", ".homefeed"})
}

func assertFileContents(t *testing.T, path, want string) {
	t.Helper()

	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("read file %s: %v", path, err)
	}

	if string(data) != want {
		t.Fatalf("file %s = %q, want %q", path, string(data), want)
	}
}

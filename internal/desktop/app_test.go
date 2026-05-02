// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package desktop_test

import (
	"os"
	"path/filepath"
	"testing"

	"homefeed/internal/desktop"
	"homefeed/internal/workspace"
)

func TestInitializeWorkspaceBridgeUsesConfiguredRoot(t *testing.T) {
	root := filepath.Join(t.TempDir(), "Homefeed")

	app := desktop.NewApp(root)
	if got := app.WorkspaceRoot(); got != root {
		t.Fatalf("workspace root = %q, want %q", got, root)
	}

	result, err := app.InitializeWorkspace()
	if err != nil {
		t.Fatalf("initialize workspace: %v", err)
	}

	if result.Layout.Root != root {
		t.Fatalf("layout root = %q, want %q", result.Layout.Root, root)
	}

	if len(result.Feeds) != 5 {
		t.Fatalf("feed count = %d, want 5", len(result.Feeds))
	}
}

func TestImportFolderBridgeCopiesIntoSelectedFeed(t *testing.T) {
	root := filepath.Join(t.TempDir(), "Homefeed")
	if _, err := workspace.Initialize(root); err != nil {
		t.Fatalf("initialize workspace: %v", err)
	}

	source := filepath.Join(t.TempDir(), "bridge-import")
	if err := os.MkdirAll(filepath.Join(source, "nested"), 0o755); err != nil {
		t.Fatalf("create source tree: %v", err)
	}
	if err := os.WriteFile(filepath.Join(source, "notes.txt"), []byte("bridge import"), 0o644); err != nil {
		t.Fatalf("write source file: %v", err)
	}

	app := desktop.NewApp(root)
	result, err := app.ImportFolder(source, "projects")
	if err != nil {
		t.Fatalf("import folder: %v", err)
	}

	wantDestination := filepath.Join(root, "projects", filepath.Base(source))
	if result.DestinationPath != wantDestination {
		t.Fatalf("destination path = %q, want %q", result.DestinationPath, wantDestination)
	}

	if _, err := os.Stat(filepath.Join(wantDestination, "notes.txt")); err != nil {
		t.Fatalf("imported file missing: %v", err)
	}
}

func TestCreateNotePostBridgeCreatesNoteFile(t *testing.T) {
	root := filepath.Join(t.TempDir(), "Homefeed")
	if _, err := workspace.Initialize(root); err != nil {
		t.Fatalf("initialize workspace: %v", err)
	}

	app := desktop.NewApp(root)
	result, err := app.CreateNotePost("projects", "Bridge Note", "bridge body")
	if err != nil {
		t.Fatalf("create note post: %v", err)
	}

	wantPath := filepath.Join(root, "projects", "bridge-note.md")
	if result.Path != wantPath {
		t.Fatalf("note path = %q, want %q", result.Path, wantPath)
	}

	if _, err := os.Stat(wantPath); err != nil {
		t.Fatalf("note file missing: %v", err)
	}
}

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package desktop_test

import (
	"path/filepath"
	"testing"

	"homefeed/internal/desktop"
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

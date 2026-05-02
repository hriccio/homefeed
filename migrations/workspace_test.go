// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package migrations

import (
	"strings"
	"testing"
)

func TestWorkspaceStatementsLoadsCanonicalArtifact(t *testing.T) {
	statements, err := WorkspaceStatements()
	if err != nil {
		t.Fatalf("load statements: %v", err)
	}

	if len(statements) != 6 {
		t.Fatalf("statement count = %d, want 6", len(statements))
	}

	if !strings.HasPrefix(statements[0], "PRAGMA foreign_keys") {
		t.Fatalf("first statement = %q, want foreign key pragma", statements[0])
	}

	if !strings.Contains(statements[len(statements)-1], "INSERT OR IGNORE INTO schema_migrations") {
		t.Fatalf("last statement = %q, want schema migration insert", statements[len(statements)-1])
	}

	foundImportBatch := false
	foundPosts := false
	for _, statement := range statements {
		if strings.Contains(statement, "CREATE TABLE IF NOT EXISTS import_batches") {
			foundImportBatch = true
		}
		if strings.Contains(statement, "CREATE TABLE IF NOT EXISTS posts") {
			foundPosts = true
		}
	}
	if !foundImportBatch {
		t.Fatal("workspace migration does not define import_batches")
	}
	if !foundPosts {
		t.Fatal("workspace migration does not define posts")
	}
}

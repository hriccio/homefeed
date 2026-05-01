// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at https://mozilla.org/MPL/2.0/.

package migrations

import (
	"embed"
	"fmt"
	"strings"
)

//go:embed 0001_create_workspace.sql
var workspaceSchema embed.FS

// WorkspaceStatements returns the canonical workspace migration statements
// from the repository-visible SQL artifact.
func WorkspaceStatements() ([]string, error) {
	raw, err := workspaceSchema.ReadFile("0001_create_workspace.sql")
	if err != nil {
		return nil, fmt.Errorf("read workspace migration: %w", err)
	}

	return splitSQLStatements(string(raw)), nil
}

func splitSQLStatements(script string) []string {
	parts := strings.Split(script, ";")
	statements := make([]string, 0, len(parts))
	for _, part := range parts {
		statement := normalizeSQLStatement(part)
		if statement == "" {
			continue
		}
		statements = append(statements, statement)
	}
	return statements
}

func normalizeSQLStatement(statement string) string {
	lines := strings.Split(statement, "\n")
	cleaned := make([]string, 0, len(lines))
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" || strings.HasPrefix(trimmed, "--") {
			continue
		}
		cleaned = append(cleaned, trimmed)
	}

	return strings.TrimSpace(strings.Join(cleaned, "\n"))
}

# Implementation Notes

## Change

`0002` Workspace migration source of truth

## Intended Path

- keep `migrations/0001_create_workspace.sql` as the canonical schema artifact
- load workspace migration statements from `migrations/workspace.go`
- remove inline schema duplication from runtime code
- preserve `InitializeWorkspace` behavior and tests

## Validation

- `go test ./...`
- `go run ./app/cli/homefeed-init -root <temp>/Homefeed`

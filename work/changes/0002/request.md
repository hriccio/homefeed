# Request: Workspace Migration Source of Truth

## Intent

Remove the migration schema duplication in the Homefeed foundation so the
workspace bootstrap has one clear source of truth.

The change should keep the current first-run behavior intact while making the
SQL migration artifact the canonical schema definition for the workspace
initializer.

## Source Evidence

- `work/changes/0001/egd.md`
- `work/changes/0001/implementation.md`
- `migrations/0001_create_workspace.sql`
- `internal/workspace/migrations.go`

## Acceptance Intent

The request should be considered satisfied when:

- workspace bootstrap still creates the same workspace layout and default feed
  rows
- the schema definition is not duplicated in multiple authoritative places
- future migration changes can be made in one place without drifting runtime
  behavior from the repository artifact


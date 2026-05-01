# Impact Analysis

## Change

`0006` Import schema source of truth

## Likely Impacted Areas

- `migrations/0001_create_workspace.sql`
- `migrations/workspace_test.go`
- `internal/imports/service.go`
- `tests/server/import_test.go`
- `internal/desktop/app.go`
- `work/changes/0005/implementation.md`

## Tensions To Watch

- the import behavior must remain unchanged
- workspace initialization must still create the import table
- the runtime import service should stop inventing schema
- tests should continue to prove the same end-user behavior

## Expected Behavioral Boundaries

- the migration artifact is the source of truth for `import_batches`
- workspace initialization applies that schema before import can run
- import batches remain recorded in SQLite
- the visible workspace layout and copy behavior stay the same

## Verification Pressure

- migration artifact includes the import table
- runtime code no longer creates the same table separately
- deterministic import test still passes
- workspace bootstrap test still passes

## Notes

This is a maintenance slice, not a feature slice. It is a direct response to the
EGD warning about schema duplication.

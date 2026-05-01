# Slice 0002 - Workspace Migration Source Of Truth

## Status

Built.

## Pack

- `polyglot_client_server`

## Runtime Targets

- `server/go`

## Architecture Mode

Local-first desktop app backend with a Go workspace initializer and SQL-backed
migration artifact.

## Discovery Scope

This slice only covers the workspace schema bootstrap path and the way it loads
its migration definition.

## Intent

Keep the existing workspace foundation behavior intact while making the
workspace schema single-sourced.

The slice should preserve:

```text
first run -> ~/Homefeed exists -> .homefeed/index.sqlite exists -> default feeds exist
```

but eliminate the duplication between runtime schema definitions and the
repository migration artifact.

## Use-Case Contract

- `InitializeWorkspace`

The use case should still create or verify the workspace root, hidden metadata
root, SQLite database, and default feed records, but it should source the schema
from the migration artifact rather than repeating it inline.

## Main Business Rules

- the workspace result must not change from slice 0001
- the migration artifact under `migrations/` is the canonical schema source
- the initializer remains idempotent
- no visible sidecar metadata files are introduced

## In Scope

- load the schema from the repository migration artifact or a small loader that
  reads it
- preserve the current default feed behavior
- preserve the current workspace and database paths
- add or update tests that prove the first slice behavior still holds
- remove the duplicated executable schema definition from runtime code

## Required Ports

- filesystem access for loading migration artifacts
- SQLite persistence
- migration runner or loader

## Initial Test Plan

- integration tests for idempotent initialization still pass
- tests prove the workspace layout and feed records are unchanged
- tests prove the initializer no longer carries a second authoritative schema
  definition inline

## Scenario Definition

1. run workspace initialization against a clean temporary root
2. confirm the workspace, metadata root, database, and default feeds are
   created exactly as before
3. rerun initialization
4. confirm the result is unchanged
5. confirm schema execution is driven from the migration artifact path

## Out Of Scope

- changes to feed semantics
- changes to workspace layout
- changes to the CLI scenario surface
- changes to the desktop shell
- import, search, AI, or UI work

## Done Criteria

- the workspace foundation still behaves the same as slice 0001
- the schema definition is no longer duplicated in runtime source
- the migration artifact remains the repository-visible source of truth
- tests continue to prove idempotent first-run initialization

## Implementation Notes

- the canonical schema lives in `migrations/0001_create_workspace.sql`
- the runtime loader is implemented in `migrations/workspace.go`
- the workspace initializer consumes the migration package rather than owning
  its own schema copy

# EGD Summary

## Change

`0006` Import schema source of truth

## Reviewed Boundary

Request: remove the runtime duplication in the import workflow so the
`import_batches` schema is defined once in the repository migration artifact
and consumed by the import use case without a second ad hoc table definition.

## Observed Behavior

- the `import_batches` table is defined in the canonical workspace migration
  artifact.
- the runtime import service no longer redeclares the `import_batches` table.
- the copy-based import workflow still copies a local source folder into the
  selected feed.
- the workflow still records import batch rows in SQLite.
- workspace initialization continues to create the schema needed by import.
- `go test ./...` passes.
- `npm run build` in `frontend/` passes.
- `go build -tags wails,production -o /tmp/homefeed-wails .` passes.

## Findings

### Note - No interactive import smoke test was run

The deterministic tests prove the schema ownership cleanup and preserve the
import behavior, but this environment did not exercise a user-selected import
through the desktop UI.

## Request-Level Judgment

The request is satisfied:

- the import schema is single-sourced
- the runtime import service no longer invents schema
- the import workflow behavior remains unchanged
- the workspace bootstrap still supports import

## Return To Loop

The current request can move toward `release`.

# EGD Summary

## Change

`0005` Local folder import workflow

## Reviewed Boundary

Request: add the first safe import workflow so the desktop app can copy a
selected local folder into the managed workspace, preserve the folder
structure, and record the import in hidden metadata without polluting visible
feed folders.

## Observed Behavior

- the app can import a local source folder through the new `ImportFolder` use
  case.
- the source folder is copied into the selected feed folder under the workspace.
- folder structure and file contents are preserved by the copy operation.
- import batch metadata is written to SQLite under the hidden `.homefeed`
  workspace database.
- visible feed folders remain free of sidecar metadata files.
- `go test ./...` passes.
- `npm run build` in `frontend/` passes.
- `go build -tags wails,production -o /tmp/homefeed-wails .` passes.

## Findings

### Warning - Import batch schema is duplicated at runtime

The import batch table exists in the repository migration artifact and is also
redeclared in `internal/imports/service.go` through `ensureImportBatchSchema`.
The slice behaves correctly, but the schema is not single-sourced yet, so future
drift remains a real maintenance risk.

### Note - No live import smoke test was run

The code and tests prove the import workflow compiles and the deterministic
scenario passes, but this environment did not exercise the new import path
through an interactive desktop session with a user-chosen folder.

## Request-Level Judgment

The request is satisfied for the first import slice:

- a local folder can be copied into a selected workspace feed
- hidden import metadata is recorded in the workspace database
- visible feed folders stay clean
- search and AI remain out of scope

## Return To Loop

The current request can move toward `release` once the maintenance tradeoff for
the duplicated schema is either accepted or refined into a follow-up slice.

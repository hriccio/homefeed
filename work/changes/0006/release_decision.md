# Release Decision

## Change

`0006` Import schema source of truth

## Decision

Accepted as the intended internal version.

## Basis

- the request asked to remove the runtime duplication in the import workflow so
  the `import_batches` schema is defined once in the repository migration
  artifact.
- the implemented slice now defines `import_batches` only in
  [migrations/0001_create_workspace.sql](/home/henrique/repos/github/hriccio/homefeed/migrations/0001_create_workspace.sql).
- the runtime import service no longer redeclares the table schema.
- `go test ./...` passes.
- `npm run build` in `frontend/` passes.
- `go build -tags wails,production -o /tmp/homefeed-wails .` passes.
- the EGD result judged the request satisfied.

## Residual Risk

- No interactive import smoke test was run in this environment, so the desktop
  UI import path was not manually exercised end to end here.

## Notes

The release boundary is the request, not just the schema cleanup slice. The
request is considered satisfied with the current internal state.

# Request: Import Schema Source Of Truth

## Intent

Remove the runtime duplication in the import workflow so the `import_batches`
schema is defined once in the repository migration artifact and consumed by the
import use case without a second ad hoc table definition.

This request keeps the import behavior itself unchanged. It only asks for the
schema to become single-sourced.

## Source Evidence

- `work/changes/0005/egd.md`
- `work/changes/0005/request.md`
- `work/changes/0005/implementation.md`
- `migrations/0001_create_workspace.sql`
- `internal/imports/service.go`
- `docs/semantics/model_hypothesis.md`
- `architecture.md`

## Acceptance Intent

The request should be considered satisfied when:

- the `import_batches` table is created by the repository migration artifact
- the import workflow no longer redeclares the same table schema at runtime
- the first import behavior continues to work exactly as before
- the workspace initialization path still creates the schema needed by import
- the deterministic import tests continue to pass

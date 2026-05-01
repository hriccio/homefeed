# Request To Slice Map

## Change

`0006` Import schema source of truth

## Mapping

| Slice | Status | Coverage | Acceptance Evidence |
| --- | --- | --- | --- |
| `docs/slices/0006-import-schema-source-of-truth.md` | built | the `import_batches` schema lives in the canonical migration artifact and the runtime import service consumes it without a duplicate table definition | deterministic tests prove workspace initialization still creates the schema and the import workflow still records batches |

## Out Of Scope

- new import features
- search indexing or search UI
- AI suggestion or mutation flows
- changes to the copy-based import behavior
- broader desktop polish

## Model Pressure

The import workflow should stay compatible with the repository's existing
single-source migration pattern. The model already treats SQLite as an
operational index, not a place to scatter schema definitions across runtime
code. This slice tightens that boundary for import metadata.

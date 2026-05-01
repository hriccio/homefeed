# Implementation Notes

## Change

`0006` Import schema source of truth

## Status

Built.

## Notes

- removed the runtime `import_batches` table redeclaration from
  `internal/imports/service.go`
- kept the copy-based import workflow and batch recording behavior unchanged
- strengthened the migration test to assert that `import_batches` exists in
  the canonical workspace migration artifact

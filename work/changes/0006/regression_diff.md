# Regression Diff

## Change

`0006` Import schema source of truth

## Result

No functional regression identified relative to the prior import workflow.

## Preserved Behavior

- the import workflow still copies a local folder into the selected feed
- import batch rows are still written to SQLite
- workspace initialization still creates the schema needed by import
- visible feed folders remain clean
- search and AI remain out of scope

## Meaningful Diff

- the `import_batches` schema is now single-sourced in the canonical workspace
  migration artifact
- the runtime import service no longer redeclares the same table schema

## Notes

This release is a maintenance cleanup of schema ownership. The user-facing
import behavior is unchanged.

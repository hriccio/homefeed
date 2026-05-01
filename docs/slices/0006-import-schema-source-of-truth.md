# Slice 0006 - Import Schema Source Of Truth

## Status

Built.

## Pack

- `polyglot_client_server`

## Runtime Targets

- `client/typescript`
- `server/go`

## Architecture Mode

Local-first desktop app with a Wails shell, a SolidJS frontend, and a Go
backend use-case layer.

## Discovery Scope

This slice only covers import schema ownership. It does not change the import
workflow behavior.

## Intent

Make the `import_batches` table single-sourced by defining it in the canonical
migration artifact and removing the runtime table redeclaration.

## Use-Case Contract

- `ImportFolder`

The import workflow should continue to copy a local folder into the selected
feed and record the batch, but it should rely on the workspace schema instead
of creating its own table definition.

## Main Business Rules

- `import_batches` is part of the workspace schema
- runtime code should not redeclare the same table schema
- the copy-based import behavior must remain unchanged
- visible feed folders remain clean
- search and AI remain out of scope

## In Scope

- move the import batch table definition into the canonical migration artifact
- remove the runtime schema redeclaration from the import service
- keep the import workflow behavior unchanged
- update tests that assert the migration statement set
- preserve the existing desktop bridge and frontend import path

## Required Ports

- SQLite persistence
- filesystem access
- Wails runtime bridge

## Initial Test Plan

- workspace initialization still creates the schema needed by import
- the import workflow still copies a local folder into the selected feed
- the import workflow still records a batch in SQLite
- the runtime service no longer defines its own duplicate schema

## Scenario Definition

1. initialize a temporary workspace
2. verify the canonical migration creates `import_batches`
3. run the existing folder import workflow
4. confirm the batch record is still written
5. confirm the visible workspace behavior is unchanged

## Out Of Scope

- new import capabilities
- import UI redesign
- search indexing
- AI behavior
- move-based import

## Done Criteria

- the import batch schema is defined only in the repository migration artifact
- import runtime code no longer redeclares that schema
- the import workflow still passes deterministic validation

## Implementation Notes

- keep the slice narrow and maintenance-focused
- preserve all user-facing import behavior
- align schema ownership with the repository's migration source-of-truth pattern
- the import batch schema now lives only in the migration artifact, and the
  runtime import service consumes it without redeclaring it

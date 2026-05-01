# Slice 0005 - Local Folder Import

## Status

Planned.

## Pack

- `polyglot_client_server`

## Runtime Targets

- `client/typescript`
- `server/go`

## Architecture Mode

Local-first desktop app with a Wails shell, a SolidJS frontend, and a Go
backend use-case layer.

## Discovery Scope

This slice covers the first safe import workflow only: copy a selected local
folder into the managed workspace and record the import in hidden metadata.

## Intent

Add a minimal import workflow that proves Homefeed can ingest an existing local
folder without breaking normal filesystem access or polluting visible feed
folders.

## Use-Case Contract

- `ImportFolder`

The application should accept a source path and a target feed, copy the source
folder into the workspace, and record the import in hidden metadata.

## Main Business Rules

- the first import strategy is copy into the workspace
- visible feed folders remain free of sidecar metadata files
- import metadata belongs under `~/Homefeed/.homefeed`
- imported content should preserve folder structure
- search and AI remain out of scope

## In Scope

- a Go import use case for local folder copy
- import batch records under the hidden metadata root
- workspace file copies that preserve the source folder structure
- bridge wiring for the import workflow
- a minimal frontend trigger or input path for the import action
- deterministic tests for the copy and metadata placement rules

## Required Ports

- filesystem access
- SQLite persistence
- Wails runtime bridge
- frontend input path

## Initial Test Plan

- a deterministic scenario imports a sample folder into a chosen feed
- the imported files appear in the workspace copy location
- import metadata is stored under `.homefeed`
- visible feed folders remain clean

## Scenario Definition

1. create a temporary source folder with nested files
2. choose a target feed in the workspace
3. run the import workflow
4. confirm the source content was copied into the workspace
5. confirm import metadata exists under `.homefeed`
6. confirm no visible sidecar metadata files were created

## Out Of Scope

- search indexing
- search UI
- AI assistance
- move-based import
- remote synchronization
- feed browsing polish

## Done Criteria

- the first import workflow can copy a local folder into the workspace
- the import writes hidden metadata instead of visible sidecar files
- the workspace initialization behavior remains unchanged
- the workflow is deterministic enough to prove with tests

## Implementation Notes

- keep the slice narrow enough to implement as a single import path
- prefer explicit use-case naming over generic file manager logic
- preserve the current startup and shell behavior while adding import

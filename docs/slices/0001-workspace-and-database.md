# Slice 0001 - Workspace And Database

## Status

Built.

## Pack

- `polyglot_client_server`

## Runtime Targets

- `client/typescript`
- `server/go`

## Architecture Mode

Local-first desktop app with a SolidJS frontend and Go domain/runtime core.

## Discovery Scope

This slice only covers the workspace and persistence foundation needed to prove
Homefeed can initialize safely on first run.

## Intent

Create the smallest useful Homefeed foundation:

```text
first run -> ~/Homefeed exists -> .homefeed/index.sqlite exists -> default feeds exist
```

This slice should establish workspace conventions and persistent structure
without implementing the full desktop UI, import flow, search, or AI behavior.

## Use-Case Contract

- `InitializeWorkspace`

The use case should create or verify the workspace root, hidden metadata root,
SQLite database, and default feed records without disturbing existing normal
files.

## Main Business Rules

- the managed workspace root is `~/Homefeed`
- hidden metadata belongs under `~/Homefeed/.homefeed`
- visible feed folders remain free of sidecar metadata files
- initialization is idempotent
- default feed records are created for the five canonical feeds
- filesystem files remain normal files outside the hidden metadata area

## In Scope

- choose the concrete project scaffold path for Wails + SolidJS + Go
- initialize `~/Homefeed`
- initialize `~/Homefeed/.homefeed`
- create `index.sqlite`
- run initial migrations
- create default feeds:
  - `professional`
  - `family`
  - `projects`
  - `personal`
  - `archive`
- expose a minimal local scenario or command that proves initialization works
- add tests around path conventions and idempotent initialization

## Required Ports

- filesystem access
- SQLite persistence
- migration runner
- deterministic ID or timestamp helpers only if the implementation needs them

## Initial Test Plan

- unit tests for workspace path resolution and hidden metadata placement
- integration tests for idempotent initialization and database creation
- a local scenario or command that exercises first-run initialization

## Scenario Definition

1. start with no `~/Homefeed`
2. run initialization
3. confirm `~/Homefeed`, `~/Homefeed/.homefeed`, and `index.sqlite` exist
4. confirm the five default feed records exist
5. run initialization again
6. confirm the result is unchanged and no visible metadata files appear beside
   feed folders

## Out Of Scope

- importing files or folders
- full feed UI
- search
- AI provider integration
- file preview
- destructive filesystem operations
- cloud sync or collaboration

## Candidate Acceptance Criteria

- running the first local scenario creates `~/Homefeed` when missing
- rerunning initialization is idempotent
- visible default feed folders are created
- hidden `.homefeed` folders are created
- SQLite database exists with initial schema
- default feed rows are present in the database
- no visible sidecar metadata files are created beside feed files

## Done Criteria

- the workspace initializes successfully on first run
- the SQLite database is created with the initial schema
- the five default feed records exist after initialization
- rerunning the initializer is safe and idempotent
- visible feed folders stay clean
- tests prove the path conventions and idempotent behavior

## First Extraction Input

- `work/sources/initial_handoff/homefeed-codex-handoff.md`

## Implementation Notes

- the first executable surface is a Go CLI scenario at `app/cli/homefeed-init`
- migrations are represented as SQL files under `migrations/`
- the default workspace path is not configurable in this slice; the CLI uses the
  current user's home directory plus `/Homefeed`

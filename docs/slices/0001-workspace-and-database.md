# Slice 0001 - Workspace And Database

## Status

Proposed.

## Intent

Create the smallest useful Homefeed foundation:

```text
first run -> ~/Homefeed exists -> .homefeed/index.sqlite exists -> default feeds exist
```

This slice should establish workspace conventions and persistent structure
without implementing the full desktop UI, import flow, search, or AI behavior.

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

## First Extraction Input

- `work/sources/initial_handoff/homefeed-codex-handoff.md`

## Open Design Decisions

- whether the first executable surface should be a Wails app command, a Go CLI scenario, or both
- whether migrations should be embedded Go files or SQL files under `migrations/`
- whether default workspace path should be configurable before the settings UI exists

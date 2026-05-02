# Slice 0007 - Create Note Post

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

This slice only covers the first note-post workflow. It does not expand into
comments, tags, search, or AI behavior.

## Intent

Add a minimal note-post workflow that creates a normal file in a selected feed
and records the post in hidden workspace metadata.

## Use-Case Contract

- `CreateNotePost`

The application should accept a feed slug, a note title, and a note body, create
the note as a normal file inside the selected feed, and store a durable post
record in SQLite-backed hidden metadata.

## Main Business Rules

- note posts are normal files in the selected feed
- post metadata lives under `~/Homefeed/.homefeed`
- visible feed folders remain free of sidecar metadata files
- the note filename should be derived deterministically from the title
- search and AI remain out of scope

## In Scope

- a Go use case for note creation
- a durable post record in SQLite
- note file creation inside a selected feed
- bridge wiring for the note workflow
- a minimal frontend trigger or input path for note creation
- deterministic tests for file creation, metadata placement, and visible-folder cleanliness

## Required Ports

- filesystem access
- SQLite persistence
- Wails runtime bridge
- frontend input path

## Initial Test Plan

- a deterministic scenario creates a note in a temporary workspace
- the note file appears in the selected feed
- the post record exists in SQLite
- no visible sidecar metadata files are created

## Scenario Definition

1. initialize a temporary workspace
2. choose a selected feed
3. create a note with a title and body
4. confirm the note file exists in the feed folder
5. confirm the post record exists in SQLite
6. confirm no visible sidecar metadata files were created

## Out Of Scope

- comments
- tags
- search indexing
- search UI
- AI assistance
- import behavior changes
- feed browsing polish

## Done Criteria

- the first note-post workflow creates a normal file in the selected feed
- the note is recorded in hidden metadata
- the workspace bootstrap and import behavior remain unchanged
- the workflow is deterministic enough to prove with tests

## Implementation Notes

- keep the slice narrow enough to implement as one post-creation path
- prefer explicit use-case naming over generic note manager logic
- preserve the current startup, import, and shell behavior while adding note
  creation

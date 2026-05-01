# Model Hypothesis

## Purpose

Homefeed models a local-first social layer over the filesystem. It treats files
and related metadata as feed posts while preserving ordinary filesystem access.

This document is the initial hypothesis. Treat it as input for `extract` and
`refine`, not final domain truth.

## Governing Sentence

Files remain normal files. Homefeed adds a social layer over the local
filesystem.

## Core Boundary

Homefeed separates:

- filesystem storage: normal files under `~/Homefeed`
- operational index: SQLite under `~/Homefeed/.homefeed/index.sqlite`
- portable metadata: hidden records under `~/Homefeed/.homefeed/meta`
- desktop experience: Wails shell with a SolidJS frontend
- AI assistance: suggestions and summaries that require user confirmation for destructive or organizing actions

Visible feed folders should remain clean. Do not create visible sidecar metadata
files beside user files.

## Candidate Concepts

- `Workspace`: the managed root at `~/Homefeed`.
- `Feed`: a top-level folder that behaves like a feed.
- `Post`: a file, note, link, folder import, AI summary, decision log, reference, or comment-thread anchor.
- `Profile`: an owner identity such as a person, group, company, project, system, or unknown source.
- `FileRef`: metadata for a normal filesystem file.
- `Comment`: local context attached to a post.
- `Tag`: a label used for filtering and discovery.
- `ImportBatch`: a tracked import operation from an external source path.
- `AgentSuggestion`: AI-generated metadata or summary that is separate from user-confirmed metadata.

## Candidate Use Cases

- initialize workspace
- create default feeds
- create a note post
- import files or folders into a feed
- assign owner and auto-create profile
- add tags and comments
- list the unified feed
- filter by feed, owner, profile, tag, file type, date, import batch, or AI classification
- search posts and metadata
- request AI metadata suggestions without applying them automatically

## Initial State Flow

Workspace setup:

```text
missing_workspace -> workspace_created -> database_created -> default_feeds_created
```

Import flow:

```text
source_selected -> batch_created -> files_copied -> file_refs_created -> posts_created -> indexed
```

AI suggestion flow:

```text
post_selected -> suggestion_requested -> suggestion_stored -> user_accepts_or_rejects
```

## Hard Rules

- files must remain normal filesystem files
- metadata must not pollute visible folders
- SQLite is the operational index, not the only source of recoverable meaning
- the first import strategy should be copy into workspace
- search returns posts, not raw files only
- AI must not move or delete files automatically in the MVP
- destructive or large organizing actions require confirmation
- the app should remain understandable if it breaks

## Open Questions

- Should the first scaffold use Wails immediately, or begin with a Go domain/core plus later Wails shell?
- Should metadata JSON export be canonical from day one or generated from SQLite?
- Should the initial UI include import flows, or only workspace/feed/post setup?
- Should the first search implementation use SQLite FTS5 immediately or plain indexed queries first?
- How much of `.homefeed/meta` should be written in the first slice?

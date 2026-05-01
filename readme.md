# Homefeed

Homefeed is a local-first social file manager for Linux desktop.

It turns normal filesystem artifacts into feed-like posts with owners, profiles,
comments, tags, import history, search, and an AI-assistance boundary. The
filesystem remains usable without the app; Homefeed must not lock files into a
proprietary database.

## Governing Sentence

Files remain normal files. Homefeed adds a social layer over the local
filesystem.

## Initial Target

- platform: Linux Mint 21.3
- app type: desktop app
- stack: Wails + SolidJS + Go + SQLite
- workspace: `~/Homefeed`
- metadata root: `~/Homefeed/.homefeed`

## Initial Sources

Original handoff material is preserved under:

- `work/sources/initial_handoff/homefeed-codex-handoff.md`

Start extraction from that file, then refine the model into repository artifacts
before implementing code.

## Initial Product Shape

Homefeed should support:

- folders as feeds
- a unified main feed
- posts that reference files, notes, links, folders, summaries, references, and comments
- owners and auto-created profiles
- comments as local context metadata
- file and folder import into the workspace
- search over posts and metadata
- an AI-agent interface that can suggest, summarize, classify, and answer without moving or deleting files automatically

## MRL Usage

This repository was created from `wastingnotime/mrl-starter`.

Read these files before substantial work:

- `AGENTS.md`
- `docs/operating/mrl_reference.md`
- `docs/operating/skills_workflow.md`
- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `docs/slices/0001-workspace-and-database.md`

Recommended next phase: run `extract` from the initial handoff source, then run
`refine` for the first minimal slice.

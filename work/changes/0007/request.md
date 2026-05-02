# Request: Create Note Post Workflow

## Intent

Add the first note-post workflow so the desktop app can create a note inside a
selected feed as a normal file while recording the post in hidden workspace
metadata.

This request focuses on the smallest social posting action after import. It does
not ask for comments, tags, search, AI, or feed browsing polish.

## Source Evidence

- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `architecture.md`
- `groundrules.md`
- `work/changes/0006/release_decision.md`
- `work/changes/0006/egd.md`

## Acceptance Intent

The request should be considered satisfied when:

- a note post can be created in a selected feed
- the note is stored as a normal file inside the workspace
- the post is recorded in hidden SQLite-backed metadata
- visible feed folders remain free of sidecar metadata files
- the workflow is deterministic enough to prove with a repeatable local
  scenario
- import, search, and AI remain out of scope

# Request To Slice Map

## Change

`0007` Create note post workflow

## Mapping

| Slice | Status | Coverage | Acceptance Evidence |
| --- | --- | --- | --- |
| `docs/slices/0007-create-note-post.md` | built | the first note-post workflow, including note file creation in a selected feed, hidden post metadata, and a minimal desktop trigger | a deterministic local scenario that creates a note and proves the file, metadata, and clean visible folder rules |

## Out Of Scope

- comments
- tags
- search indexing or search UI
- AI suggestion or mutation flows
- import behavior changes
- feed browsing polish

## Model Pressure

The model wants Homefeed to treat posts as the primary social unit while
preserving normal filesystem usability. This slice should make a note a normal
file in the selected feed and keep the durable post record in hidden metadata
instead of inventing visible sidecar files.

# Impact Analysis

## Change

`0007` Create note post workflow

## Likely Impacted Areas

- `internal/posts/`
- `migrations/`
- `internal/desktop/`
- `frontend/`
- `tests/`
- `docs/semantics/` if the note-post contract needs a small clarification

## Tensions To Watch

- note files must remain ordinary files on disk
- post metadata must stay hidden under `.homefeed`
- the first note workflow should stay small and deterministic
- the new feature should not drag in comments, tags, search, or AI
- filename generation must avoid hidden randomness

## Expected Behavioral Boundaries

- the user can create a note in a selected feed
- the note appears as a normal file in the feed folder
- a durable post record exists in SQLite
- visible feed folders stay clean
- the workspace initialization and import behavior remain unchanged

## Verification Pressure

- note file exists with expected content
- post metadata exists in SQLite
- no visible sidecar metadata files appear beside the note
- repeatable local scenario

## Notes

This slice is the first step toward treating posts as the main user-facing unit
without broadening into the rest of the social graph.

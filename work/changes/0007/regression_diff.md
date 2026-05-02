# Regression Diff

## Change

`0007` Create note post workflow

## Result

No functional regression identified relative to the prior workspace and import
flows.

## Preserved Behavior

- workspace initialization still creates the same workspace layout and default
  feed rows
- import still copies local folders into the selected feed and records batches
  in SQLite
- visible feed folders remain clean
- search and AI remain out of scope

## Meaningful Diff

- the desktop shell now exposes a `CreateNotePost` bridge
- the frontend now includes a note creation form
- the system now records note posts as normal files plus hidden SQLite-backed
  metadata

## Notes

This release adds the first note-post workflow while preserving the workspace
and import behavior established by earlier slices.

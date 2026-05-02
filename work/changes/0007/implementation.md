# Implementation Notes

## Change

`0007` Create note post workflow

## Status

Built.

## Notes

- added a Go `CreateNotePost` use case in `internal/posts`
- wrote note posts as normal markdown files inside the selected feed
- recorded note post metadata in SQLite through the workspace database
- wired the note workflow through the desktop bridge and SolidJS frontend
- added deterministic tests for the note file, post record, and bridge path

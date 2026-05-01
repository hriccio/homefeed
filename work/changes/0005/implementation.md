# Implementation Notes

## Change

`0005` Local folder import workflow

## Status

Built.

## Notes

- added a Go `ImportFolder` use case in `internal/imports`
- stored import batch records in SQLite under the hidden workspace metadata
  root
- wired the import workflow through the desktop bridge and SolidJS frontend
- added deterministic tests for folder copying, batch metadata, and bridge
  wiring

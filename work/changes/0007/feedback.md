# Feedback

## Change

`0007` Create note post workflow

## Evidence

- Date: 2026-05-01
- Source: live desktop smoke test follow-up from the user
- Symptom: the UI briefly showed a message like `no bridge available`

## Observed Result

The note-post workflow itself appeared to work, but the SolidJS shell could
report that the Wails bridge was unavailable while the window was still
attaching. That made the startup state noisier than necessary.

## Interpretation

This looks like a bridge-readiness race in the frontend rather than a failure of
the note-post use case.

## Next Loop Impact

The frontend should wait for the Wails bindings before enabling the workspace,
import, and note actions so the user sees a stable ready state instead of a
transient bridge warning.

# Feedback

## Change

`0007` Create note post workflow

## Evidence

- Date: 2026-05-01
- Source: live desktop smoke test follow-up from the user
- Symptom: the UI showed a transient `no bridge available` style message and a
  waiting state below the initialization button
- Symptom: disabled buttons used an hourglass-style hover cue
- Symptom: the `Initialize workspace` button could be blocked with `not
  allowed`
- Symptom: clicking `Initialize workspace` immediately after startup could
  still report that the Wails bridge was unavailable

## Observed Result

The note-post workflow itself still appeared to work, but the shell made the
startup state look noisier than necessary.

## Interpretation

The issue is a bridge-attachment race in the click path, not a broken note-post
use case.

## Next Loop Impact

The shell should remain quiet while the Wails bridge is attaching, and actions
should wait briefly before failing so the user can click through startup
without hitting a false unavailable message.

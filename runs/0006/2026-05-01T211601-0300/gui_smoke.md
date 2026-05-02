# GUI Smoke Test

## Change

`0006` Import schema source of truth

## Timestamp

`2026-05-01T211601-0300`

## Environment

- Host: local Linux desktop session
- WebKit development package: installed
- Build tags: `wails,production`

## Steps

1. Built the production Wails binary with `go build -tags wails,production -o /tmp/homefeed-wails .`
2. Launched `/tmp/homefeed-wails` in the desktop session.
3. Located the live `Homefeed` window with `wmctrl -lp`.
4. Queried the window geometry with `xwininfo`.
5. Focused the window and clicked the `Initialize workspace` button using X11 events.
6. Confirmed the window was still present after the click.

## Result

- A top-level `Homefeed` window appeared.
- The `Initialize workspace` button received a click in the live GUI.
- The window stayed open after the click.
- The launch log only contained the WebKit signal warning:
  - `Overriding existing handler for signal 10. Set JSC_SIGNAL_FOR_GC if you want WebKit to use a different signal`

## Interpretation

The desktop shell passes an interactive GUI smoke test in this environment and
responds to a basic click-through interaction.

# GUI Smoke Test

## Change

`0004` SolidJS frontend for the desktop shell

## Timestamp

`2026-05-01T201325-0300`

## Environment

- Host: local Linux desktop session
- WebKit development package: installed
- Build tags: `wails,production`

## Steps

1. Built the production Wails binary with `go build -tags wails,production -o /tmp/homefeed-wails .`
2. Launched `/tmp/homefeed-wails` in the desktop session.
3. Inspected the window list with `wmctrl -lp`.

## Result

- A top-level `Homefeed` window appeared.
- The launch log only contained the WebKit signal warning:
  - `Overriding existing handler for signal 10. Set JSC_SIGNAL_FOR_GC if you want WebKit to use a different signal`

## Interpretation

The desktop shell now passes an interactive GUI smoke test in this environment.

# Implementation Notes

## Change

`0003` Minimal Wails desktop shell

## Intended Path

- add the first Wails shell around the existing initializer
- preserve the current workspace bootstrap implementation
- keep the frontend surface minimal
- place the Wails entrypoint at the repository root behind the `wails` build
  tag
- keep the default build path stubbed so tests remain runnable in headless
  environments

## Validation

- desktop launch smoke check
- `go test ./...`
- `go build -tags wails .`

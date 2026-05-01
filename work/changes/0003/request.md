# Request: Minimal Wails Desktop Shell

## Intent

Expose the existing Homefeed workspace initializer through a minimal Wails
desktop shell so the product begins to match the selected desktop runtime shape.

The change should keep the current workspace bootstrap behavior intact while
adding the first desktop entrypoint and a thin Go-to-frontend bridge.

## Source Evidence

- `work/changes/0001/request.md`
- `work/changes/0001/egd.md`
- `work/changes/0002/request.md`
- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `architecture.md`

## Acceptance Intent

The request should be considered satisfied when:

- the application has a Wails desktop shell entrypoint
- the shell can invoke the existing workspace initializer through a thin bridge
- the first desktop scenario still creates the same workspace layout and default
  feed rows
- import, search, and AI behavior remain out of scope

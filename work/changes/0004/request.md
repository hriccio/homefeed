# Request: SolidJS Frontend For The Desktop Shell

## Intent

Replace the static HTML/JavaScript shell assets with a real SolidJS frontend
while keeping the existing Wails desktop entrypoint and workspace bridge
behavior intact.

The request is to make the frontend runtime shape match the selected
architecture without introducing import, search, AI, or broader app features.

## Source Evidence

- `work/changes/0003/request.md`
- `work/changes/0003/egd.md`
- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `architecture.md`
- `docs/packs/polyglot_client_server.md`

## Acceptance Intent

The request should be considered satisfied when:

- the desktop shell uses a SolidJS frontend build
- the frontend can invoke the existing workspace initializer through the thin
  Go bridge
- the workspace bootstrap behavior remains unchanged
- the built frontend still remains minimal and focused on initialization proof


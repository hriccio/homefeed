# Repository Guidelines

## Project Structure & Module Organization
This repository is an MRL-based project named Homefeed. Strategic docs live at the root and working design material lives under `docs/`. Read `docs/operating/mrl_reference.md` and `docs/operating/skills_workflow.md` before changing the workflow.

Project-specific context lives in:

- `readme.md`
- `docs/semantics/model_hypothesis.md`
- `docs/semantics/domain_background_knowledge.md`
- `docs/slices/0001-workspace-and-database.md`
- `work/sources/initial_handoff/`

On the first pass through this repository's guidance, consider `.agents/skills/adoption-diagnose/SKILL.md` before substantial project-specific work. Use it when licensing, README content, selected pack, semantic placeholders, or other starter-adoption decisions have not clearly been settled in repository artifacts.

The selected implementation shape is Wails + SolidJS + Go + SQLite. Use this structure as the target shape after scaffold:

```text
app/
frontend/
internal/{domain,workspace,feeds,profiles,posts,files,imports,index,search,comments,tags,agent,jobs,config,logging}
migrations/
tests/
docs/{operating,building,evaluation,semantics,slices}
.agents/skills/
```

Record structural deviations in `decisions.md`.

## Build, Test, and Development Commands
Keep tooling lightweight until the first slice exists.

- `wails dev` runs the desktop app after scaffold exists.
- `go test ./...` runs Go tests after Go modules exist.
- `npm test` or the package-manager equivalent runs frontend tests after the SolidJS app exists.
- a small Go initialization scenario is acceptable before the full Wails shell exists.

## Coding Style & Naming Conventions
Prefer Go for backend/domain behavior, SolidJS for UI, explicit types, and business-oriented names. Use verb-driven use cases such as `InitializeWorkspace` and intention-revealing repositories such as `GetByID` and `Save`.

## Testing Guidelines
Use tests as specification. Start with domain tests, add integration tests for mappings and end-to-end flows, and keep time, IDs, and external responses deterministic.

## Commit & Pull Request Guidelines
Use Conventional Commits for commit subjects, choosing an appropriate type such as `feat`, `fix`, `docs`, `refactor`, `test`, `build`, `ci`, or `chore`. Commit after every completed and verified change before starting unrelated work. Keep commits scoped to one request, slice, or doc change, and include test evidence in pull requests.

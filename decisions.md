# Decisions Log

## Purpose

This document records relevant architectural and implementation decisions for the adopting repository.

Use it to preserve reasoning, avoid re-discussing settled trade-offs without context, and document deviations from `architecture.md` and `groundrules.md`.

---

## Entry Template

```md
## DEC-XXXX - Title

- Date: YYYY-MM-DD
- Status: proposed | accepted | superseded | rejected
- Owners: human | codex | both

### Context
What problem or tension led to this decision?

### Decision
What was decided?

### Consequences
What becomes easier, harder, or different because of this?

### Alternatives considered
What other options were considered and why were they not chosen?

### Notes
Any additional implementation guidance, migration note, or follow-up.
```

---

## Index

Add entries as the repository evolves.

## DEC-0013 - Adopt Homefeed As The Project Identity

- Date: 2026-05-01
- Status: accepted
- Owners: both

### Context
This repository was created from `wastingnotime/mrl-starter` to support a
local-first social file manager for Linux desktop.

The source context selects Wails + SolidJS + Go + SQLite and defines the core
product boundary: the filesystem remains usable without the app, while Homefeed
adds feed, post, profile, comment, import, search, and AI-assistance layers.

### Decision
The adopting repository is named `homefeed`.

It will model a local-first desktop app with:

- normal files under `~/Homefeed`
- hidden metadata and SQLite under `~/Homefeed/.homefeed`
- folders as feeds
- posts as the primary user-facing unit
- owners/profiles and comments as local context
- AI suggestions separated from user-confirmed metadata

The repository adopts the `polyglot_client_server` pack because the selected
implementation has a SolidJS frontend and Go/Wails backend runtime.

### Consequences
The repository now has project-specific README, semantic, source, and first
slice artifacts. Future Codex sessions should treat the copied starter material
as workflow infrastructure, not as the Homefeed domain.

The first implementation should establish workspace and database initialization
before import, search, and AI behavior.

### Alternatives considered
Use the starter's default `python_ddd_monolith` pack. This was rejected because
the chosen runtime is Wails + SolidJS + Go + SQLite.

Start with full Wails UI implementation immediately. This was deferred because
the first slice needs the filesystem and persistence foundation before broad UI
work.

### Notes
The original handoff file is preserved under `work/sources/initial_handoff/`.

## DEC-0001 - Separate MRL Core From Implementation Packs

- Date: 2026-03-29
- Status: accepted
- Owners: both

### Context
The starter was presenting Python plus a DDD-inspired modular monolith as if that were the default shape of MRL itself. That creates confusion when a repository needs another language, another architecture such as event sourcing, or more than one runtime.

### Decision
The repository now distinguishes between:

- MRL core, which stays artifact-driven and architecture-agnostic
- implementation packs, which define language, architecture, structure, and testing defaults

This repository selects `polyglot_client_server` for Homefeed. Other adopting
repositories may choose different packs.

### Consequences
It becomes easier to reuse the same refinement workflow across Python, JavaScript, Go, event-sourced, and polyglot client/server repositories. It also becomes necessary to make the selected pack explicit in architecture docs and slice docs.

### Alternatives considered
Keep one universal Python starter and treat every other shape as an undocumented deviation. This was rejected because it would keep conflating MRL with one implementation style.

### Notes
Future pack additions should live under `docs/packs/` and should be referenced by slice documents when the runtime topology matters.

## DEC-0002 - Treat Skill Model Guidance As Advisory

- Date: 2026-04-02
- Status: accepted
- Owners: both

### Context
The repository skills now include model guidance for tasks such as `build`, `refine`, and `extract`. That creates a potential ambiguity: a reader could assume that naming a preferred model in a skill will automatically switch the active Codex model or force sub-agent routing during execution.

### Decision
Model guidance inside repository skills is advisory only. It documents which model shape is usually a good fit for the task, but it does not by itself require automatic model switching, worker spawning, or hard routing behavior.

### Consequences
The skills remain durable even if model names, availability, or routing capabilities change. The repository gains clearer guidance for future operators and tooling, but predictable model selection still requires explicit runtime policy or orchestration outside the skill text.

### Alternatives considered
Encode specific model names in skills as if they were enforced execution rules. This was rejected because skill text alone does not guarantee runtime behavior and would overstate what the repository can currently control.

### Notes
If the repository later wants deterministic skill-to-model routing, document that as a separate operational decision and implement it in the calling workflow or agent orchestration layer.

## DEC-0003 - Avoid `.codex` Repository Artifacts For Now

- Date: 2026-04-02
- Status: accepted
- Owners: both

### Context
The repository has used `.codex`-specific artifacts while shaping the workflow. That creates a risk that the MRL process starts to look tool-defined rather than process-defined before it is clear whether MRL can stay tool-agnostic in practice.

### Decision
For now, the repository should avoid relying on `.codex` as part of the committed process shape. The workflow should stay centered on MRL artifacts and repository documents rather than tool-specific folders. This decision is explicitly reversible while the team validates whether MRL can remain tool-agnostic or whether a specific AI tool will prove operationally necessary.

### Consequences
The repository stays cleaner and more focused on portable process artifacts. It may require some extra translation when using Codex or another tool because fewer tool-native conventions are captured directly in versioned files. The decision also preserves room to adopt tool-specific support later if real usage shows that generic artifacts are not enough.

### Alternatives considered
Keep `.codex` as a first-class part of the repository workflow immediately. This was rejected for now because it would prematurely optimize for one tool before validating whether that coupling is necessary for MRL.

### Notes
If future evidence shows that MRL depends on stable capabilities from a specific AI tool, record a follow-up decision describing the required coupling, why generic artifacts were insufficient, and which tool-specific assets should become part of the repository.

## DEC-0004 - Separate Expose Extensions From Packs

- Date: 2026-04-16
- Status: accepted
- Owners: both

### Context
Some adopting repositories will repeatedly expose accepted artifacts through the same remote infrastructure path, such as GitHub Actions, container publication, and a handoff to a separate infrastructure repository. That recurring shape should be reusable, but it does not belong in MRL core and does not cleanly fit the repository's definition of an implementation pack.

### Decision
The repository distinguishes expose extensions from implementation packs.

Expose extensions are repository-local lifecycle guides for recurring exposure mechanisms. They define how accepted released artifacts are exposed into a real target context, how handoff boundaries are described, and which evidence artifacts should be recorded.

Implementation packs remain responsible for implementation defaults such as language, architecture, structure, testing layout, and runtime topology.

The starter provides generic guidance in `docs/operating/expose_extensions.md`, a concrete example extension in `docs/operating/extensions/expose_aws_ecr_infra_pr.md`, and a reusable per-change template in `work/changes/_template/exposure.md`.

### Consequences
Repositories can standardize recurring exposure mechanisms without redefining the MRL loop and without overloading packs with lifecycle behavior. Adopting repositories must still record their chosen expose extension and repository-specific exposure boundary in their own decisions and change artifacts.

### Alternatives considered
Treat every exposure path as ad hoc per change, or encode remote deployment behavior as a pack. These were rejected because recurring exposure paths deserve explicit reuse while remaining separate from implementation defaults.

### Notes
An adopting repository that uses one expose extension as its normal path should add its own accepted decision naming that extension and defining what counts as exposure completion for that repository.

## DEC-0005 - Treat Consumer-Facing Interfaces As Exposed Contracts

- Date: 2026-04-16
- Status: accepted
- Owners: both

### Context
Some repositories or repository families expose APIs, emitted events, webhooks, or other integration interfaces that downstream consumers need to discover and rely on across releases. A runtime artifact alone is not enough when consumers also need to know which interfaces exist, which payloads are offered, and whether a change is breaking, additive, or internal-only.

### Decision
The starter recognizes exposed contracts as a repository-level architectural policy for consumer-facing interfaces.

When a repository adopts this policy:

- externally consumable interfaces are treated as part of the exposed surface
- tracked contract artifacts should live in explicit repository locations
- contract artifacts should be produced or updated during `build`
- contract compatibility should be reviewed during `release`
- exposure evidence should reference the released contract artifacts during `expose`

This guidance is captured in `docs/operating/exposed_contracts.md`.

### Consequences
Repositories that expose consumer-facing interfaces can make those surfaces discoverable and governable across releases. This improves compatibility review and downstream reuse, but it also creates an obligation to keep contract artifacts synchronized with implementation and release evidence.

### Alternatives considered
Treat APIs and events as discoverable only through code, runtime inspection, or informal documentation after exposure. This was rejected because consumers need a clearer and more durable contract surface than reverse engineering can provide.

### Notes
This is not a mandatory rule for every MRL repository. It is intended for repositories or repository families whose released artifacts are expected to be consumed through explicit interfaces.

## DEC-0006 - Treat Key Operational Telemetry As An Exposed Surface

- Date: 2026-04-16
- Status: accepted
- Owners: both

### Context
Some repositories or repository families are expected to be operated through metrics, health signals, traces, and structured logs. If those signals are only discovered after deployment inside an APM or logging platform, the operational surface becomes accidental and hard to review. Metrics are especially contract-like when dashboards, alerts, SLOs, autoscaling, or external analysis depend on specific names, labels, and meanings.

### Decision
The starter recognizes exposed observability as a repository-level architectural policy for operationally relied-on telemetry.

When a repository adopts this policy:

- key operational telemetry is treated as part of the exposed surface
- observability artifacts should be documented in explicit repository locations
- key signal definitions should be produced or updated during `build`
- meaningful observability changes should be reviewed during `release`
- exposure evidence should reference the operational signal locations during `expose`

This guidance is captured in `docs/operating/exposed_observability.md`.

### Consequences
Repositories can make their operational surface more intentional before exposure instead of relying on post-deploy discovery alone. This improves operator clarity and compatibility review, but it also creates an obligation to keep important telemetry definitions synchronized with implementation and release evidence.

### Alternatives considered
Treat observability as a runtime-only concern discovered inside an APM after deployment. This was rejected because important telemetry surfaces, especially metrics and health signals, are often stable enough to deserve explicit review and documentation.

### Notes
This is not a mandatory rule for every MRL repository. It is intended for repositories or repository families whose running systems are expected to be monitored and consumed through stable operational signals.

## DEC-0007 - Commit Completed Changes Immediately

- Date: 2026-04-27
- Status: accepted
- Owners: both

### Context
MRL treats repository artifacts as durable memory and execution contexts as stateless. Leaving completed work only in the active working tree makes the workflow depend on hidden operator state and weakens handoff between isolated phases.

### Decision
After every completed and verified repository change, the operator should create a focused Conventional Commit before starting unrelated work. This applies to workflow artifacts, semantic artifacts, slice documents, code, tests, and release or exposure evidence.

### Consequences
Repository history becomes part of the MRL memory model. Isolated phase runs can trust committed state more easily, but operators must avoid batching unrelated changes just because they occurred in the same session.

### Alternatives considered
Rely on end-of-session commits or leave commit timing to each operator. This was rejected because it preserves too much hidden state between completed changes.

### Notes
If a change must remain uncommitted temporarily, record why in the active change artifact or tell the operator explicitly.

## DEC-0008 - Run EGD Against Requests

- Date: 2026-04-27
- Status: accepted
- Owners: both

### Context
The workflow described EGD as if the slice were the primary review boundary. Slices are useful implementation increments, but a user request can be broader than one slice or can require checking whether several slice-level decisions add up to the intended change.

### Decision
EGD should run against the request. Slice documents, implementation artifacts, tests, and scenario packets are evidence for the request-level review, not the boundary of the review itself. Release should likewise decide whether the request has been satisfied well enough to accept internally.

### Consequences
Expectation-gap review is less likely to miss omissions that sit between slices. EGD artifacts and run folders should be keyed by request or change id when practical. Build and refine can remain slice-oriented because implementation still benefits from bounded increments.

### Alternatives considered
Keep EGD slice-oriented and rely on release to detect request-level gaps. This was rejected because EGD is the phase explicitly responsible for finding expectation gaps before acceptance.

### Notes
A request may still map to a single slice. In that case, request-level EGD should be lightweight and can use the one slice as its main evidence packet.

## DEC-0009 - Track Request-To-Slice Mapping During Refine

- Date: 2026-04-27
- Status: accepted
- Owners: both

### Context
After making EGD request-centered, the workflow needs an explicit handoff from request intent to slice execution. Without a durable map, build can remain slice-local while EGD and release have to reconstruct how those slices were supposed to satisfy the request.

### Decision
`refine` should maintain `work/changes/<id>/request_slice_map.md` whenever a request artifact exists. The map records whether the request is covered by one slice or multiple slices, what each slice covers, what is out of scope, and which evidence should support request-level EGD and release.

### Consequences
Build can stay bounded by slice while still preserving traceability back to request intent. EGD and release gain a direct artifact for checking whether the built slices add up to the requested change. Refinement has one more artifact to maintain, but it removes ambiguity at the phase boundary.

### Alternatives considered
Store this mapping only inside each slice document or rely on `impact_analysis.md`. This was rejected because the request-to-slice relationship is a cross-slice concern and deserves a stable per-change artifact.

### Notes
For a one-slice request, the map should still exist but can be short.

## DEC-0010 - Release Requests, Not Slices

- Date: 2026-04-27
- Status: accepted
- Owners: both

### Context
EGD now evaluates request satisfaction and `refine` records request-to-slice traceability. Release must use the same boundary or it can accept a slice that is locally complete while the request remains incomplete.

### Decision
`release` accepts, rejects, or returns the implemented request state. Slice documents and implementations are evidence for the release decision, but they are not the release boundary.

### Consequences
Release decisions align with request-level EGD and the request-to-slice map. Exposure and feedback should refer to the accepted request state, while still preserving links to the supporting slices and implementation artifacts.

### Alternatives considered
Keep release slice-oriented and rely on EGD to report request-level gaps. This was rejected because acceptance should happen at the same boundary as expectation-gap review.

### Notes
When a request maps to exactly one slice, the release evidence may be short but should still name the request being accepted.

## DEC-0011 - Add Adoption Diagnosis And Guidance Skills

- Date: 2026-04-27
- Status: accepted
- Owners: both

### Context
Adopting repositories need help distinguishing MRL starter material from project-specific decisions. Two recurring needs do not fit the canonical loop phases: auditing whether starter defaults have been intentionally replaced, and helping an owner understand MRL workflow choices without changing files.

### Decision
Add two support skills outside the canonical loop:

- `adoption-diagnose`, which audits adoption readiness around licensing, README content, pack selection, project-specific artifacts, and starter residue
- `guidance`, which answers owner/operator questions about MRL phases, artifacts, skills, and boundaries from repository evidence without mutating files

### Consequences
The starter gains a clearer adoption path for repository owners while keeping the canonical loop unchanged. Support skills must remain role-bounded and artifact-driven so they do not become hidden workflow state or broad autonomous agents.

### Alternatives considered
Fold these responsibilities into `extract` or `refine`. This was rejected because adoption readiness and owner guidance are operational support concerns rather than semantic extraction or slice design.

### Notes
`adoption-diagnose` may write `work/adoption_diagnosis.md` when a durable report is useful. `guidance` is read-only by default.

## DEC-0012 - Surface Adoption Diagnosis At Repository Entry

- Date: 2026-04-27
- Status: accepted
- Owners: both

### Context
Some adoption decisions should happen early, before substantial project-specific work makes starter defaults harder to distinguish from intentional choices. `AGENTS.md` is the first repository guidance many Codex sessions read, so it is the right place to surface that check.

### Decision
`AGENTS.md` should instruct agents to consider `adoption-diagnose` on the first pass through repository guidance before substantial project-specific work, especially when licensing, README content, selected pack, semantic placeholders, or other starter-adoption decisions are not clearly settled in artifacts.

### Consequences
Agents get an early reminder to check adoption readiness without making `adoption-diagnose` part of every loop phase. The check remains conditional: once the repository artifacts clearly show those decisions are settled, normal phase work can proceed without repeating the audit.

### Alternatives considered
Require `adoption-diagnose` before every skill run. This was rejected because it would add noise after adoption decisions are already explicit.

### Notes
This is an entry guardrail, not a substitute for owner decisions.

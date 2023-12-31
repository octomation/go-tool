---
code: null
id: I_kwDODa89W88AAAABSsOR_A
databaseId: 5549298172
number: 556
url: https://github.com/octomation/go-tool/issues/556
title: "ci/cd: release policy and curated changelog notes"
issueType: null
labels:
  - "type: feature"
  - "scope: docs"
  - "scope: inventory"
  - "impact: medium"
  - "effort: medium"
assignees:
  - kamilsk
milestone: null
parent: null
issueFields: {}
state: OPEN
stateReason: null
duplicateOf: null
createdAt: 2026-09-23T05:40:25Z
updatedAt: 2026-09-23T13:24:30Z
lastEditedAt: null
closedAt: null
---

# ci/cd: release policy and curated changelog notes

**Motivation:** a release is published by pushing a `v*` tag, and nothing checks that the release was prepared. For some projects the goreleaser changelog built from commits is enough; others keep hand-written notes (e.g. `docs/content/changelog/<tag>.md` in octopot/indexit) and want them as the release body, with a failed push, not a failed release, when the note is missing. The template should make this a conscious, per-project choice rather than an accident.

**Release policy to choose from:**

- `generated` (default) — goreleaser changelog from commits, no checks; zero effort, low-quality notes.
- `curated, optional` — use the note if it exists, fall back to `generated`; flexible, but a forgotten note goes unnoticed.
- `curated, required` — no note, no release; best notes, costs writing one for every tag, including patch releases.

**Convention to agree on:**

- where notes live and how a tag maps to a file;
- versioned git hooks in `.github/hooks` (no new top-level folders), wired by `make setup` via `core.hooksPath`; reconcile with the current `bin/lib/git/hooks` + `GIT_HOOKS` mechanism;
- a `pre-push` hook refusing a `v*` tag whose commit lacks the note (local, bypassable with `--no-verify`);
- the same check as the first step of `cd.yml`, so a bypassed hook fails in seconds, before the build;
- turning a docs page into a release body: frontmatter, the duplicated `# title`, site-relative links;
- how this fits docs and changelog publishing, so one note serves the site and the GitHub release.

**PoC:** octopot/indexit — `.github/hooks/pre-push`, the "Check the changelog note" step in `.github/workflows/cd.yml`.

**Out of scope:** signing commits from goreleaser ([[issue-20]]).

<!-- 2026-09-23T07:08Z https://github.com/octomation/go-tool/issues/556#issuecomment-5790588582
Hook wiring (`.github/hooks` via `core.hooksPath`, and what to do with `bin/lib/git/hooks` + `GIT_HOOKS`) is tracked in octomation/makefiles#81.
-->

<!-- 2026-09-23T13:24Z https://github.com/octomation/go-tool/issues/556#issuecomment-5795657731
Settled design, implemented in octopot/indexit (uncommitted PoC):

- **One entry point:** `.github/scripts/release.mjs`, zero-dependency Node, with modes `check` (offline, used by the `pre-push` hook, `make release-check` and the first step of `cd.yml`), `render` (note → release body: strips frontmatter and the H1, makes site-relative Markdown/HTML links absolute from `SITE_URL`, prints the title), `preflight` (CI: required secrets are set and usable) and `doctor` (on demand, with remediation).
- **Note convention:** `docs/content/changelog/<tag>.md`, `.md` only, frontmatter `title`/`description`, first line `# <title>`, HTML allowed, MDX rejected. The note must be committed in the tagged commit and match the worktree.
- **Release branch:** the tag commit must be on its branch, pushed in the same push or already on `origin/<branch>`: the first `x-release.branches` rule (`{match: "^v(\\d+)\\.", branch: "v$1"}` for maintenance lines) or the default branch. `git push --atomic origin <branch> <tag>`.
- **Config:** `.github/settings.json` stores only what cannot be derived (release policy, required secrets); `.github/settings.cue` documents it; `cue` comes from the tools module. Owner/name, default branch, site URL (`configure-pages` → `SITE_URL` for Nextra and the renderer) and the tap (`.goreleaser.yml`) are derived. Probot's `settings.yml` stays; an own settings diff/apply is deferred.
- **goreleaser:** `release.name_template: '{{ envOrDefault "RELEASE_TITLE" .Tag }}'`, `release.preflight.fail_on_error: true` (needs ≥ v2.18), no hardcoded `release.github`; release with `github.token`, tap with its own secret.
- **Policies:** `curated, required` is what the hook enforces; `generated` stays possible by not installing the hook and dropping the check step.

Make targets: octomation/makefiles#82 (`release-check`), #83 (`doctor`), #84 (`config-vet`); hooks: octomation/makefiles#81.
-->

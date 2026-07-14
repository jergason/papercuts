# papercuts

Coding agents often run across busted workflows or small frictions caused by a repository. They're getting better at working around them, but these papercuts still cost tokens and time, and might trip up humans too!

This is a tiny CLI for noting repo-owned papercuts so we can tell OTHER agents to go back and fix them later without breaking flow. A papercut belongs when a plausible fix would change code, tests, scripts, configuration, or documentation in the current repository. I love the metaphor of "sanding down" the rough edges on a codebase. This should help!

This is a reimplementation of the technique [Steve Ruiz described on X](https://x.com/steveruizok/status/2075303919664734295): a `papercuts` CLI plus an `AGENTS.md` instruction telling agents to log friction proactively, and a periodic pass where a model fixes whatever's accumulated. All credit for the idea goes to him - this repo just packages it up as an installable CLI and skill.

## What's here

- **`papercuts`** - a small, dependency-free Go binary. `papercuts -m <model> "message"` appends a timestamped entry to `PAPERCUTS.md` at the root of the current git repo.
- **`skills/papercuts-fix`** - an [agent skill](https://github.com/vercel-labs/skills) that sweeps `PAPERCUTS.md`, prunes misrouted entries, fixes what's safely fixable, and leaves substantial repo-owned work for a human to decide.
- **`agents-md-snippet.md`** - a copy-pasteable instruction block for your `AGENTS.md` / `CLAUDE.md` telling agents to log papercuts proactively.

## Install the CLI

```sh
# read this first, we're in the wild west
curl -fsSL https://raw.githubusercontent.com/jergason/papercuts/main/install.sh | sh
```

Downloads the right prebuilt binary for your OS/arch from the [latest release](https://github.com/jergason/papercuts/releases/latest), verifies its checksum, and installs it to `$HOME/.local/bin` (override with `INSTALL_DIR`). No Go, Node, or other toolchain required on the machine that runs it.

Or, if you already have Go: `go install github.com/jergason/papercuts/cmd/papercuts@latest`

## Usage

```sh
papercuts -m gpt-5-codex 'The document-management Jest suite passes but leaves an open handle. Update its teardown or shared test harness so the worker exits cleanly.'
```

Appends to `PAPERCUTS.md` at the repo root, found by walking up from `cwd` looking for `.git`.

```
2026-07-10T18:15:58.874Z - gpt-5-codex - Jamison Dance

The document-management Jest suite passes but leaves an open handle. Update its teardown or
shared test harness so the worker exits cleanly.
```

The `-m` flag and message are both required. Author is auto-detected from `git config user.name`, falling back to `$USER`.

## Set up proactive logging

Copy the block from [`agents-md-snippet.md`](./agents-md-snippet.md) into your project's `AGENTS.md` or `CLAUDE.md`. That's the instruction that gets agents to actually use this thing without being asked.

Generic shell mistakes, agent or tool limitations, sandbox restrictions, external service behavior, and friction owned by another repository do not belong in the current repository's log. If the repository should detect, document, or accommodate an external problem, describe that repo-local deficiency instead.

## Install the fix-it skill

```sh
npx skills add jergason/papercuts --skill papercuts-fix
```

Then, every so often, ask your agent to fix the papercuts (or trigger however your skill runner invokes it). It'll read `PAPERCUTS.md`, prune clearly out-of-scope entries, fix the small repo-owned stuff, and leave substantial or ambiguous repo-owned work for you.

## Building from source

```sh
go build ./cmd/papercuts
```

Releases are cut with [GoReleaser](https://goreleaser.com/) via `.github/workflows/release.yml` on tag push (`vX.Y.Z`).

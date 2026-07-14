---
name: papercuts-fix
description: Sweep this repo's PAPERCUTS.md and fix the small logged frictions that are safe to fix. Use when the user asks to fix papercuts, clean up papercuts, address logged papercuts, or "sand down" the repo.
disable-model-invocation: true
---

# Fix papercuts

`PAPERCUTS.md` is a running log of small, repo-owned frictions. A valid entry has a plausible fix in a version-controlled file in this repository: code, tests, scripts, configuration, or documentation.

## Steps

1. Read `PAPERCUTS.md` at the repo root. If it doesn't exist or is empty, say so and stop.
2. Classify each entry before changing code:
   - **Repo-owned:** A plausible fix changes a version-controlled file in this repository.
   - **Out of scope:** The friction is a generic shell mistake, agent or tool limitation, sandbox restriction, external service behavior, or belongs to another repository.
   - Treat an external problem as repo-owned only when this repository should detect, document, or accommodate it.
3. Remove clearly out-of-scope entries from `PAPERCUTS.md`. Record their disposition for the final summary, including the owning repository when apparent.
4. For each repo-owned entry, decide if it is **safely fixable in isolation**: a typo or stale command in docs, a missing or wrong script, a broken link, a small config fix, or an obviously wrong path. Make the fix.
5. Leave repo-owned entries that are ambiguous, require a design decision, or touch behavior beyond a small local correction. Do not guess.
6. Remove fixed entries from `PAPERCUTS.md`. Leave its header and unresolved repo-owned entries intact.
7. Report what was fixed, what out-of-scope material was pruned or rerouted, and what repo-owned work remains. Do not commit anything.

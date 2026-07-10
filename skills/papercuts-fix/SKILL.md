---
name: papercuts-fix
description: Sweep this repo's PAPERCUTS.md and fix the small logged frictions that are safe to fix. Use when the user asks to fix papercuts, clean up papercuts, address logged papercuts, or "sand down" the repo.
disable-model-invocation: true
---

# Fix papercuts

`PAPERCUTS.md` is a running log of small frictions other agents hit while working in this repo: dead-end tool calls, broken links, misleading errors, undocumented setup steps. Individually none of these were worth stopping for, but together they're worth cleaning up.

## Steps

1. Read `PAPERCUTS.md` at the repo root. If it doesn't exist or is empty, say so and stop.
2. Go entry by entry. Each entry has a `timestamp - model - author` header line followed by a short description of the friction (and sometimes a guessed cause/fix).
3. For each entry, decide if it's **safely fixable in isolation**: a typo or stale command in docs, a missing/wrong script, a broken link, a small config fix, an obviously wrong path. Make the fix.
4. Skip entries that are ambiguous, require a design decision, or touch behavior beyond a small local correction. Leave those in `PAPERCUTS.md` untouched - don't guess.
5. Once an entry is fixed, remove that entry (its header line and message) from `PAPERCUTS.md`. Leave the file's `# Papercuts` header and any remaining entries intact.
6. When done, report a short summary: what was fixed, and what was left behind and why. Don't commit anything - that's the user's call.

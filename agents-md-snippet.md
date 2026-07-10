Paste this into your project's `AGENTS.md` or `CLAUDE.md` so robots log papercuts while working.

```markdown
# Log papercuts

Important! When you hit a small friction while working - a tool call that missed and had to be retried, a confusing or undocumented setup step, a flaky command, a stale cache, a misleading error, a non-obvious gotcha - log it to `PAPERCUTS.md` via `papercut -m <model> "message"`. One or two sentences: what you were doing -> what got in the way (a guess at the cause/fix is a bonus). Do this proactively, in the moment, even though none of these are blocking. Logged together they show where the repo needs sanding down. This is distinct from a changelog and from tracked issues (real bugs / tracked work).
```

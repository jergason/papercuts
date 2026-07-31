Paste this into your project's `AGENTS.md` or `CLAUDE.md` so robots log papercuts while working.

```markdown
# Log repo papercuts

When you hit a small friction caused by this repository that could be fixed by a change to a version-controlled file (code, tests, scripts, config, docs), log it to `PAPERCUTS.md` via `papercuts -m <model> 'message'`.

In one or two sentences, name the affected command, path, or subsystem and describe the likely repo-local improvement when apparent.

Do not log generic shell mistakes, agent or tool limitations, sandbox restrictions, external service behavior, or friction owned by another repository. If this repository should detect, document, or accommodate an external problem, log that repo-local deficiency instead.
```

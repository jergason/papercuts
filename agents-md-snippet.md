Paste this into your project's `AGENTS.md` or `CLAUDE.md` so robots log papercuts while working.

```markdown
# Log repo papercuts

When you hit a small friction caused by this repository, log it to `PAPERCUTS.md` via `papercuts -m <model> 'message'`.

Log it only when a plausible fix would change a version-controlled file in this repository: code, tests, scripts, configuration, or documentation. In one or two sentences, name the affected command, path, or subsystem and describe the likely repo-local improvement when apparent.

Do not log generic shell mistakes, agent or tool limitations, sandbox restrictions, external service behavior, or friction owned by another repository. If this repository should detect, document, or accommodate an external problem, log that repo-local deficiency instead. This log is also distinct from a changelog and from tracked issues.
```

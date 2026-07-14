# Papercuts

Small frictions logged by agents while working in this repo: dead-end tool calls, broken links, misleading errors, undocumented setup steps. Not accomplishments (see LOG.md) and not tracked bugs (see the issue tracker) - just the sand in the gears. Logged via `papercuts`.

2026-07-10T21:49:35.139Z - jamison-not-a-robot - Jamison Dance

The install script failed at first I think?

2026-07-10T21:56:33.430Z - claude-sonnet-5 - Jamison Dance

Testing a GH Actions env-var-driven shell command locally via `VAR=val uvx "tool@${VAR}"` silently failed (empty expansion) because inline prefix assignments aren't visible to same-line $VAR expansion in bash the way env: blocks are in Actions; had to export first. Minor gotcha when dry-running workflow run: steps locally.

2026-07-14T14:51:43.423Z - gpt-5-codex - Jamison Dance

Running go test ./... in the managed workspace failed because Go's default build cache under ~/Library/Caches was read-only. Setting GOCACHE to a writable /tmp directory avoids the sandbox-specific failure.

2026-07-14T14:52:27.321Z - gpt-5-codex - Jamison Dance

Creating an annotated git tag failed with 'unable to create temporary file' because this managed workspace allowed git commit but blocked the tag object's write inside .git. The tag operation needed an escalated retry.


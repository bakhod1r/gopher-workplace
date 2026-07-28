# Recognized vs default

## The idea

Return two bools: the parsed value and whether the input was recognized. That
lets the caller distinguish an explicit `false` from an unknown string that
should keep a default. Missing an accepted form (`off`) silently turns it into
"unknown".

## Why it matters

Config and env parsing must recognize every documented form. An omitted case
means a valid setting is treated as absent, and the default silently wins — a
subtle production misconfiguration.

## Watch out

- Normalize (`ToLower` + `TrimSpace`) before matching.
- Keep truthy and falsey sets symmetric (`on`/`off`, `yes`/`no`).
- Unknown input returns `ok=false`, not `false,true`.

# Publish A Snapshot Without Tearing It

**Level:** staff
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A config reloader updates two fields of a shared struct in place. Under load a request occasionally sees the new retry count with the old timeout, and the combination is not one anybody configured.

## Task

Implement [atomicsnapshot.go](atomicsnapshot.go):

1. Publish `c` as the current snapshot.
2. Readers must never observe a half-updated Config.
3. The store must not alias the caller's variable — later writes to it are invisible.

Replace the stub body in [atomicsnapshot.go](atomicsnapshot.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s.Set(Config{Retries:3, Timeout:10}); s.Get()
Output: {3 10}
```

**Example 2:**

```
Input:  c := Config{...}; s.Set(c); c.Retries = 99; s.Get()
Output: the value as of Set
```

_Explanation:_ The store holds its own copy.

**Example 3:**

```
Input:  Get before any Set
Output: the zero Config
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Atomic pointer swap** | Publishing one word replaces a whole struct indivisibly. |
| 2 | **Immutable snapshots** | Never mutate a published value; publish a new one. |
| 3 | **Escape by design** | The snapshot must be heap-allocated — it outlives the call and is shared. |
| 4 | **Happens-before** | `Store`/`Load` on an atomic gives readers the fully written struct. |

## Hint

Two fields cannot be written atomically. One pointer can.

## Validate

```bash
make verify
```

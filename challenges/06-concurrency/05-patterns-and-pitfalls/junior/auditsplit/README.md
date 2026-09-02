# Tee Audit Events

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Every audit event has to reach two places: the long-term archive and the
real-time alerting rules. Reading the source stream twice is not an option —
a channel value goes to exactly one receiver — so the stream is *teed* into
two channels that each get a full copy.

## Task

Implement `TeeAudit` in [auditsplit.go](auditsplit.go) so that:

1. It creates two channels and returns them, with one goroutine feeding both.
2. For each event, it sends a copy to both channels, in whichever order they become ready.
3. Both channels are closed once `events` is drained.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TeeAudit(chan of "login")
Output: both channels yield "login", then both close
```

**Example 2:**

```
Input:  TeeAudit(chan of login, update, logout)
Output: both channels yield all three in order
```

**Example 3:**

```
Input:  TeeAudit(closed empty channel)
Output: both channels close immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Tee pattern** | One source, two independent consumers, each seeing every value. |
| 2 | **Nil channel in select** | Setting a case's channel to `nil` disables it, so each copy is sent once. |
| 3 | **Independent consumers** | Both outputs must be read concurrently or the tee blocks. |

## Hint

Copy the two channels into local variables, then run a two-iteration select
that sets each one to `nil` after it has taken its copy.

## Validate

```bash
make verify
```

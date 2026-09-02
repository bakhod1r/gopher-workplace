# Draining a Notification Buffer

**Level:** junior
**Topic:** 06-concurrency → 04-context-and-cancellation

## Context

A websocket subscriber drains the notifications buffered for its user and writes them to the socket in one batch. Sockets die constantly — laptops sleep, tunnels drop — so the drain must stop on cancellation and hand back what it managed to collect, letting the caller requeue the rest for the next connection.

## Task

Implement the exported function(s) in [drainnotifications.go](drainnotifications.go) so that:

1. It loops on a `select` over `ctx.Done()` and a comma-ok receive from `ch`.
2. It appends each value, preserving arrival order.
3. A closed channel returns the collected slice and `nil`.
4. A finished context returns the partial slice and `ctx.Err()`, never `nil`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Drain(live ctx, closed chan "a","b")
Output: ["a", "b"], nil
```

**Example 2:**

```
Input:  Drain(live ctx, closed empty chan)
Output: [], nil
```

**Example 3:**

```
Input:  Drain(cancelled ctx, empty chan)
Output: [], context.Canceled
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Accumulate-then-return** | The partial result is data, not garbage — the caller requeues it. |
| 2 | **Comma-ok receive** | Only `ok == false` means the producer is finished. |
| 3 | **Non-nil empty slice** | `make([]string, 0)` differs from `nil` under `reflect.DeepEqual`. |

## Hint

`got := make([]string, 0)` up front, then `for { select { ... } }`, returning `got` on both exits.

## Validate

```bash
make verify
```

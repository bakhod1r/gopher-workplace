# Actor Pattern

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An actor owns its state outright: no other goroutine touches `count`. Callers
send *work* down a channel, and a single goroutine executes it serially. The
mutual exclusion comes from the channel, not from a lock.

## Task

Implement `Add` on `*CounterActor` in [actorpatt.go](actorpatt.go):

1. Send a `func(*int)` on `a.msgs`.
2. That closure must add `n` to the int it is given.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  New(); Add(1); drain
Output: count == 1
```

**Example 2:**

```
Input:  100 goroutines each calling Add(1)
Output: count == 100, no data race under -race
```

**Example 3:**

```
Input:  Add(5); Add(-2)
Output: count == 3 (messages applied in send order)
```

_Explanation:_ `run` processes the channel one message at a time.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Serialized state** | Only `run` dereferences `&a.count`, so there is exactly one writer. |
| 2 | **Behaviour as a message** | The channel element is a `func(*int)` — the actor receives operations, not data. |
| 3 | **Closure capture** | The closure captures `n` from the `Add` call, which is what makes it self-contained. |

## Hint

One statement: `a.msgs <- func(c *int) { *c += n }`. Do not touch `a.count`
directly in `Add` — that is the race the actor exists to remove.

## Validate

```bash
make verify
```

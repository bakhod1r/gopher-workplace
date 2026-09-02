# Drain Delivery Queue

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The webhook delivery worker is shutting down. Before it exits it must empty
its inbox so the dispatcher's pending sends complete, and record how many
undelivered attempts were discarded.

## Task

Implement `DrainQueue` in [queuedrain.go](queuedrain.go) so that:

1. It receives from `attempts` until the dispatcher closes it.
2. It ignores the attempts themselves and returns how many there were.
3. An already-closed empty queue returns `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DrainQueue(1, 2, 3)
Output: 3
```

**Example 2:**

```
Input:  DrainQueue() // closed, empty
Output: 0
```

**Example 3:**

```
Input:  DrainQueue(9)
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`for range ch`** | Ranging without a variable when the value is irrelevant. |
| 2 | **Draining** | Emptying a channel so blocked senders can finish. |
| 3 | **Counter** | The only state a drain loop needs. |

## Hint

`for range attempts { ... }` — with no variable — is legal and idiomatic
when you only care that a value arrived.

## Validate

```bash
make verify
```

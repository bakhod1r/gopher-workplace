# Future / Promise

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A future is a value that is not ready yet. Readers block on it; one writer
completes it. The synchronization primitive for "this has happened, and it can
only happen once" is a closed channel — every receive on a closed channel
returns immediately, for any number of waiters.

## Task

Implement `Complete` and `Get` on `*Future` in [promisefut.go](promisefut.go):

1. `Complete(val)` stores `val` in `f.val` and then closes `f.done`.
2. `Get()` blocks by receiving from `f.done`, then returns `f.val`.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Get() while another goroutine Completes(42) after 10ms
Output: blocks ~10ms, returns 42
```

**Example 2:**

```
Input:  Complete(7); then Get() three times
Output: 7, 7, 7   (Get is repeatable)
```

**Example 3:**

```
Input:  50 goroutines in Get(); one Complete(99)
Output: all 50 return 99
```

_Explanation:_ closing a channel releases every waiter at once — a send would release only one.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Close as a broadcast** | `close(ch)` unblocks all current and future receives. |
| 2 | **Write-then-close ordering** | The close is the happens-before edge that publishes `f.val` safely. |
| 3 | **Non-blocking check with `select`/`default`** | How `IsDone` polls without waiting. |

## Hint

Store the value **before** closing. Closing first is a data race: a waiter can
wake and read `f.val` while `Complete` is still writing it.

## Validate

```bash
make verify
```

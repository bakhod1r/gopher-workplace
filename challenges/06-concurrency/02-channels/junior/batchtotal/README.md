# Batch Total

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The nightly settlement batch stages the day's transaction amounts on a
buffered channel so the same totalling code runs whether the amounts came
from a file or from the ledger service.

## Task

Implement `BatchTotal` in [batchtotal.go](batchtotal.go) so that:

1. It stages every amount on one channel with capacity `len(amounts)`.
2. It closes the channel before reading it back.
3. It returns the settlement total using a `range` loop; an empty day returns `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  BatchTotal([]int{100, 250, 99})
Output: 449
```

**Example 2:**

```
Input:  BatchTotal(nil)
Output: 0
```

**Example 3:**

```
Input:  BatchTotal([]int{-200, 200})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Send, close, range** | The single-goroutine channel idiom. |
| 2 | **Capacity sizing** | `len(amounts)` guarantees no send blocks. |
| 3 | **Termination by close** | `range` needs the close to ever return. |

## Hint

Order matters: all sends, then `close`, then `range`. Closing first would
make the sends panic.

## Validate

```bash
make verify
```

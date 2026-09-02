# Line Totals

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The order-pricing service runs two stages end to end: the first converts
unit counts to cents at two cents a unit, the second adds a one-cent
handling fee. Each stage is a goroutine reading one channel and writing the
next.

## Task

Implement `LineTotals` in [feepipeline.go](feepipeline.go) so that:

1. Stage 1 emits the unit counts on a channel and closes it.
2. Stage 2 receives them, sends `n * 2` cents on a second channel, and closes that.
3. The collector receives from stage 2 and adds the one-cent fee, so each count becomes `n*2+1`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  LineTotals([]int{1, 2})
Output: [3 5]
```

**Example 2:**

```
Input:  LineTotals([]int{0})
Output: [1]
```

**Example 3:**

```
Input:  LineTotals(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Pipeline** | Stages linked by channels, each closing its own output. |
| 2 | **Close propagation** | Stage 1's close ends stage 2's range, which triggers stage 2's close. |
| 3 | **Order preservation** | A single-goroutine chain keeps line order exactly. |

## Hint

Each stage owns exactly one channel: it ranges over its input and closes
its output when that range ends.

## Validate

```bash
make verify
```

# Invoice Total

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The nightly billing run totals an invoice's line items in cents. The lines
are spread across a few worker goroutines that share one queue; each worker
adds up what it pulled and the coordinator adds up the workers.

## Task

Implement `TotalCents` in [invoicetotal.go](invoicetotal.go) so that:

1. It puts every line on a shared queue channel and closes it.
2. It starts `workers` goroutines, each ranging over the queue and reporting a partial sum.
3. It waits for all workers, then returns the invoice total. `workers < 1` behaves as `1`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TotalCents([]int{100, 250, 99}, 2)
Output: 449
```

**Example 2:**

```
Input:  TotalCents(nil, 4)
Output: 0
```

**Example 3:**

```
Input:  TotalCents([]int{500}, 1)
Output: 500
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fan-out** | Many goroutines `range` over one channel; each line goes to exactly one. |
| 2 | **`sync.WaitGroup`** | `Add` before the goroutine, `defer Done` inside, `Wait` before closing partials. |
| 3 | **Closing partials** | Close only after `wg.Wait()`, or the collecting `range` ends early. |

## Hint

Close the queue before starting the workers so their `range` loops
terminate. Call `wg.Wait()` before `close(partials)`.

## Validate

```bash
make verify
```

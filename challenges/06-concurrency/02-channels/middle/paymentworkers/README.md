# Payment Capture Workers

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

A checkout service captures authorised payments through a provider that allows only a few concurrent calls. A fixed pool of workers pulls charge IDs off a jobs channel and pushes outcomes onto a results channel; the caller needs every outcome, keyed by charge, no matter which worker handled it.

## Task

Implement the stubbed function in [paymentworkers.go](paymentworkers.go) so that:

1. Start `workers` goroutines that pull charges off a jobs channel.
2. Collect every outcome into a map keyed by charge ID.
3. Call `capture` exactly once per charge and never exceed `workers` concurrent calls.
4. A `workers` count of zero or less means one worker.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CaptureAll([]string{"ch_1"}, 2, ok)
Output: map[ch_1:captured]
```

**Example 2:**

```
Input:  CaptureAll([]string{"ch_1","ch_2"}, 1, ok)
Output: map[ch_1:captured ch_2:captured]
```

**Example 3:**

```
Input:  CaptureAll(nil, 4, ok)
Output: map[]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | Worker pool | Fixed goroutine count, unbounded job list — the pool size, not the input size, sets the concurrency. |
| 2 | Close ownership | The feeder closes `jobs`; a separate goroutine closes `results` after `wg.Wait()`. |
| 3 | Collect on the main goroutine | Ranging over `results` in the caller means the map is written by exactly one goroutine — no mutex needed. |
| 4 | Deadlock avoidance | `wg.Wait()` must not run on the goroutine that also drains `results`, or the workers block forever on their sends. |

## Hint

Three goroutine roles: workers, one feeder that closes `jobs`, and one closer that does `wg.Wait()` then `close(results)`. The caller just ranges over `results`.

## Validate

```bash
make verify
```

# Spam Filter

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A mail filter screens an inbound batch against a banned phrase. Messages are
screened concurrently, and the filter returns one verdict per message rather
than a filtered list, so the caller can keep every message aligned with its
original envelope.

## Task

Implement `Flagged` in [spamfilter.go](spamfilter.go) so that:

1. Return a `[]bool` the same length as `messages`.
2. Element `i` reports whether `messages[i]` contains `banned`.
3. Screen each message in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Flagged([]string{"buy now", "hello"}, "buy")
Output: [true false]
```

**Example 2:**

```
Input:  Flagged([]string{"hello"}, "buy")
Output: [false]
```

**Example 3:**

```
Input:  Flagged(nil, "buy")
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Mark, do not filter** | Producing one verdict per slot avoids the shared state a concurrent filter would need. |

## Hint

Build a parallel array of verdicts and let the caller filter sequentially.
Collecting matches from goroutines would need shared state and lose the order.

## Validate

```bash
make verify
```

# Cache Warmer

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A read-through cache is warmed on deploy so the first real users do not pay the
cold-start cost. Every key is loaded from the origin concurrently, and the
warmer records how many bytes each entry occupies — never more than the per-
entry cap, and never negative when the origin reports a miss.

## Task

Implement `WarmAll` in [cachewarmer.go](cachewarmer.go) so that:

1. Return a slice of sizes the same length as `keys`.
2. Element `i` is `load(keys[i])`, raised to `0` when negative and lowered to `capBytes` when larger.
3. Load each key in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  WarmAll([]string{"a", "bb"}, loader, 1000)
Output: [100 200]
```

**Example 2:**

```
Input:  WarmAll([]string{"huge"}, loader, 150)
Output: [150]
```

**Example 3:**

```
Input:  WarmAll(nil, loader, 100)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Shared read-only parameters** | `load` and `capBytes` are read by every goroutine and written by none — always safe. |

## Hint

Clamp inside the goroutine using a local variable, then write `out[i]` exactly
once. Two `if` statements are enough.

## Validate

```bash
make verify
```

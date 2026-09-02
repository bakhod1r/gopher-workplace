# Rate Limiter

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

An API gateway renders a quota dashboard for all tenants on a plan. For each
tenant it subtracts the requests already spent this window from the plan limit,
clamping at zero so an over-limit tenant never shows a negative allowance.
Tenants are computed concurrently.

## Task

Implement `RemainingQuota` in [ratelimiter.go](ratelimiter.go) so that:

1. Return a slice of remaining quotas the same length as `used`.
2. Quota `i` is `limit - used[i]`, raised to `0` when negative.
3. Compute each tenant in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RemainingQuota([]int{10, 90}, 100)
Output: [90 10]
```

**Example 2:**

```
Input:  RemainingQuota([]int{150}, 100)
Output: [0]
```

**Example 3:**

```
Input:  RemainingQuota(nil, 100)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Shared read-only limit** | `limit` is read by every goroutine and written by none, which is always safe. |

## Hint

Compute into a local variable inside the goroutine, clamp it, and then write
`out[i]` exactly once.

## Validate

```bash
make verify
```

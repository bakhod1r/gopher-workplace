# Cert Checker

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A certificate monitor runs nightly over every TLS certificate the fleet serves.
Each certificate's expiry is stored as a day number, and the monitor flags the
ones that expire within the alert window — including those already expired.
Certificates are checked concurrently.

## Task

Implement `ExpiringSoon` in [certchecker.go](certchecker.go) so that:

1. Return a `[]bool` the same length as `expiries`.
2. Flag `i` is `true` when `expiries[i] - today <= window`, so an already-expired certificate is flagged too.
3. Check each certificate in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ExpiringSoon([]int{100, 400}, 90, 30)
Output: [true false]
```

**Example 2:**

```
Input:  ExpiringSoon([]int{50}, 90, 30)
Output: [true]
```

**Example 3:**

```
Input:  ExpiringSoon(nil, 90, 30)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Two shared read-only inputs** | `today` and `window` are captured by every goroutine and modified by none. |

## Hint

Compute the days remaining into a local variable first — it makes the boundary
condition (`daysLeft == window`) easy to read and easy to get right.

## Validate

```bash
make verify
```

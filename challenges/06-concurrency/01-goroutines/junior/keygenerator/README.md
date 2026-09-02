# Key Generator

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A key-generation service draws a pool of random candidates and must find which
of them are prime before building a key. Testing one candidate says nothing
about the others, so the whole pool is tested concurrently and each result is
recorded against its candidate.

## Task

Implement `PrimeCandidates` in [keygenerator.go](keygenerator.go) so that:

1. Return a `[]bool` the same length as `candidates`.
2. Element `i` reports whether `candidates[i]` is prime; anything below `2` is not prime.
3. Test each candidate in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PrimeCandidates([]int{2, 4, 7})
Output: [true false true]
```

**Example 2:**

```
Input:  PrimeCandidates([]int{1, 0, -5})
Output: [false false false]
```

**Example 3:**

```
Input:  PrimeCandidates(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **CPU-bound fan-out** | Goroutines are not only for I/O; independent CPU work spreads across cores the same way. |

## Hint

Trial-divide while `d*d <= n`. Seeding the verdict with `n >= 2` makes small and
negative candidates fall out for free.

## Validate

```bash
make verify
```

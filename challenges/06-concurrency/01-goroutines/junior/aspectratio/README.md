# Aspect Ratio

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A video player reduces each stream's resolution to its aspect ratio for the
quality menu: 1920×1080 should read as 16:9. The reduction needs the greatest
common divisor of width and height, and every stream in the manifest is reduced
concurrently.

## Task

Implement `Divisors` in [aspectratio.go](aspectratio.go) so that:

1. Return a slice the same length as `sizes`.
2. Element `i` is the greatest common divisor of `sizes[i][0]` and `sizes[i][1]`, using absolute values.
3. Compute each divisor in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Divisors([][2]int{{1920, 1080}})
Output: [120]
```

**Example 2:**

```
Input:  Divisors([][2]int{{0, 720}})
Output: [720]
```

**Example 3:**

```
Input:  Divisors(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Arrays are values** | A `[2]int` is copied when passed to the goroutine, so each one owns its pair outright. |

## Hint

Euclid's algorithm: repeat `a, b = b, a%b` until `b` is zero, then `a` is the
answer. Take absolute values before the loop starts.

## Validate

```bash
make verify
```

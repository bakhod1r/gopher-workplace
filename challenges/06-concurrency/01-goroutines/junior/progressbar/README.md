# Progress Bar

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A CLI shows a progress bar for every job in a batch. Each bar is `width`
characters wide and filled in proportion to that job's completion percentage.
Bars are rendered concurrently and printed in job order.

## Task

Implement `Bars` in [progressbar.go](progressbar.go) so that:

1. Return a slice of bars the same length as `percents`.
2. Clamp each percentage into `[0, 100]`, then fill `pct * width / 100` characters with `#`.
3. Render each bar in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Bars([]int{50, 100}, 10)
Output: [##### ##########]
```

**Example 2:**

```
Input:  Bars([]int{0}, 10)
Output: []
```

**Example 3:**

```
Input:  Bars(nil, 10)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Guarding a panic** | `strings.Repeat` panics on a negative count, and a panic in *any* goroutine kills the whole process. |

## Hint

Clamp before calling `strings.Repeat`. A panic inside a goroutine cannot be
recovered by the caller — it takes the program down.

## Validate

```bash
make verify
```

# Disk Reporter

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A capacity report scans every mounted volume and reports the single largest file
it found, so operators can spot the runaway log. Volumes are scanned
concurrently because a slow disk must not hold up the others, and an empty
volume still gets a row.

## Task

Implement `LargestFiles` in [diskreporter.go](diskreporter.go) so that:

1. Return a slice with one size per volume, in volume order.
2. Size `i` is the largest value in `volumes[i]`; an empty volume reports `0`.
3. Scan each volume in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  LargestFiles([][]int{{30, 10}, {5}})
Output: [30 5]
```

**Example 2:**

```
Input:  LargestFiles([][]int{{}})
Output: [0]
```

**Example 3:**

```
Input:  LargestFiles(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **`defer` covers early returns** | The empty-volume branch returns early, and `defer wg.Done()` still counts it down. |

## Hint

Seed the running maximum with `vol[0]`, not `0` — but check the length first, or
the index panics.

## Validate

```bash
make verify
```

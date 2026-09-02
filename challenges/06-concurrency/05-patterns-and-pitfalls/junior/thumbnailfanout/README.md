# Thumbnail Fan Out

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Thumbnail rendering is the slowest step in the media service, and a single
renderer cannot keep up with a burst of uploads. *Fanning out* puts several
identical renderer goroutines on the same job channel: whichever renderer is
free takes the next image, so the burst spreads across all of them.

## Task

Implement `RenderThumbnails` in [thumbnailfanout.go](thumbnailfanout.go) so that:

1. It creates a jobs channel and starts exactly `workers` renderer goroutines that each range over it.
2. Every renderer sends `render(img)` to a results channel.
3. After all images are queued and every renderer has finished, the results are collected and returned sorted ascending.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RenderThumbnails([]string{"a", "b", "c"}, 2, render)
Output: []string{"a.thumb", "b.thumb", "c.thumb"}
```

**Example 2:**

```
Input:  RenderThumbnails([]string{"z"}, 4, render)
Output: []string{"z.thumb"}
```

**Example 3:**

```
Input:  RenderThumbnails(nil, 2, render)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fan-out** | N workers ranging over one channel share the work with no dispatcher of your own. |
| 2 | **WaitGroup** | `wg.Wait()` is how you know every renderer returned before closing results. |
| 3 | **Order independence** | Thumbnails finish in unpredictable order, so the result must be sorted. |

## Hint

Buffer the results channel with `len(images)` so renderers never block, then
close `jobs`, `wg.Wait()`, close `results`, and drain it.

## Validate

```bash
make verify
```

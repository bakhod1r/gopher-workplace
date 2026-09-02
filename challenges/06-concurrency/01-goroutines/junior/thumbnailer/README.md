# Thumbnailer

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A media pipeline generates thumbnails for an upload batch. Before any pixels are
touched, it computes the output height of each image so that every thumbnail is
exactly `maxWidth` pixels wide and keeps its aspect ratio. Images are
independent, so each is measured in its own goroutine.

## Task

Implement `TargetHeights` in [thumbnailer.go](thumbnailer.go) so that:

1. Return a slice of heights the same length as `images`.
2. Height `i` is `images[i].Height * maxWidth / images[i].Width`.
3. An image with a non-positive width yields `0`.
4. Compute each height in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TargetHeights([]Image{{100, 50}, {200, 200}}, 100)
Output: [50 100]
```

**Example 2:**

```
Input:  TargetHeights([]Image{{0, 40}}, 100)
Output: [0]
```

**Example 3:**

```
Input:  TargetHeights(nil, 100)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Structs are copied** | Passing an `Image` into the goroutine hands it a private copy, not a shared pointer. |

## Hint

Return early for a degenerate width. Because `wg.Done()` is deferred, the early
return still releases the WaitGroup.

## Validate

```bash
make verify
```

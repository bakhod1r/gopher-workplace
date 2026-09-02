# Thumbnail Areas

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The image resize service pre-computes how many pixels each square thumbnail
will occupy, given its side length. Requests move through a channel so the
same code serves both the batch and the on-demand path.

## Task

Implement `ThumbAreas` in [thumbnailer.go](thumbnailer.go) so that:

1. It moves the values through a channel rather than transforming the slice directly.
2. Each side length becomes its squared pixel area.
3. Areas come back in request order; `nil` input yields an empty, non-nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ThumbAreas([]int{64, 128})
Output: [4096 16384]
```

**Example 2:**

```
Input:  ThumbAreas([]int{2})
Output: [4]
```

**Example 3:**

```
Input:  ThumbAreas(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Buffered channel as a queue** | Capacity `len(sides)` means no send blocks. |
| 2 | **FIFO ordering** | A channel preserves the order values were sent in. |
| 3 | **Close then range** | Close before draining when a single goroutine does both. |

## Hint

With one goroutine doing both the sending and the receiving, the buffer
must be big enough for everything — and you must `close` before you `range`.

## Validate

```bash
make verify
```

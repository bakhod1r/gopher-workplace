# Log Shipper

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A log agent batches lines before shipping them to the collector and must know
how many bytes each line will occupy on the wire — the line itself plus the
trailing newline the protocol appends. Lines are measured concurrently so sizing
never becomes the bottleneck in the hot path.

## Task

Implement `PayloadSizes` in [logshipper.go](logshipper.go) so that:

1. Return a slice of sizes the same length as `lines`.
2. Size `i` is `len(lines[i]) + 1` (bytes on the wire, newline included).
3. Measure each line in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PayloadSizes([]string{"ok", "boom"})
Output: [3 5]
```

**Example 2:**

```
Input:  PayloadSizes([]string{""})
Output: [1]
```

**Example 3:**

```
Input:  PayloadSizes(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **`len` counts bytes** | Wire size is a byte count, so `len` is correct here and rune counting would under-report. |

## Hint

The goroutine body is a single write: `out[i] = len(line) + 1`. Everything else
is the WaitGroup ceremony.

## Validate

```bash
make verify
```

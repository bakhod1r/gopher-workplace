# Video Transcoder

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A transcoding farm prepares a lower-quality rendition of every track in an
upload. Each track's target bitrate is a fixed percentage of its source bitrate,
and the tracks are sized concurrently before the encoders are started.

## Task

Implement `TargetBitrates` in [videotranscoder.go](videotranscoder.go) so that:

1. Return a slice of target bitrates the same length as `bitrates`.
2. Target `i` is `bitrates[i] * factorPct / 100`.
3. A `factorPct` of zero or less yields `0`.
4. Compute each target in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TargetBitrates([]int{4000, 2000}, 50)
Output: [2000 1000]
```

**Example 2:**

```
Input:  TargetBitrates([]int{4000}, 100)
Output: [4000]
```

**Example 3:**

```
Input:  TargetBitrates(nil, 50)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **One shared read-only setting** | `factorPct` is fixed before the loop, so every goroutine may read it without synchronisation. |

## Hint

Multiply before dividing so integer truncation does not eat the result, and
guard the non-positive factor with an early `return`.

## Validate

```bash
make verify
```

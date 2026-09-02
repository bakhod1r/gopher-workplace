# Backup Uploader

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A multipart backup uploader splits an archive into parts and must send a
checksum alongside each one so the server can verify it. Checksums are computed
concurrently while the parts are being staged, and the result must stay in part
order.

## Task

Implement `PartChecksums` in [backupuploader.go](backupuploader.go) so that:

1. Return a slice of checksums the same length as `parts`.
2. The checksum starts at `0` and folds each byte with `h = h*31 + int(b)`; an empty part checksums to `0`.
3. Checksum each part in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PartChecksums([]string{"a", "b"})
Output: [97 98]
```

**Example 2:**

```
Input:  PartChecksums([]string{"ab"})
Output: [3105]
```

**Example 3:**

```
Input:  PartChecksums(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Deterministic under concurrency** | The fold depends only on its input, so running N of them at once still gives byte-identical results. |

## Hint

Iterate over `[]byte(part)` so the fold is over bytes, not runes, and declare
the accumulator `h` inside the goroutine.

## Validate

```bash
make verify
```

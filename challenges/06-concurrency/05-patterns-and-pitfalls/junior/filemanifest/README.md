# File Manifest Sizes

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Building the manifest for a release archive means stating every file. The
stats are independent, so they run concurrently — but the manifest lines must
stay in the declared order, or the checksum of the manifest itself changes on
every build.

## Task

Implement `FileSizes` in [filemanifest.go](filemanifest.go) so that:

1. It preallocates the result slice with `make([]int, len(paths))`.
2. It starts one goroutine per path, tracked by a `WaitGroup`.
3. Each goroutine writes `size(path)` into its own index; the slice is returned after `wg.Wait()`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FileSizes([]string{"a", "bb"}, size)
Output: []int{1, 2}
```

**Example 2:**

```
Input:  FileSizes([]string{"aaaa", "b", "cc"}, size)
Output: []int{4, 1, 2}
```

**Example 3:**

```
Input:  FileSizes(nil, size)
Output: []int{} (empty, length 0)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **WaitGroup** | `Add` before starting the goroutine, `Done` deferred inside it, `Wait` before reading. |
| 2 | **Index-keyed writes** | Separate slice elements are separate memory — no lock required. |
| 3 | **Deterministic output** | Concurrency changes the timing, not the order of the result. |

## Hint

Pass `i` and `path` into the goroutine as arguments and write to `sizes[i]` —
no mutex, no channel, no sorting.

## Validate

```bash
make verify
```

# Search Indexer

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A search backend builds an inverted index. For one term it needs the term
frequency in every document of a shard, and documents are independent, so the
shard is scanned concurrently with one goroutine per document.

## Task

Implement `TermCounts` in [searchindexer.go](searchindexer.go) so that:

1. Return a slice of counts the same length as `docs`.
2. Count `i` is the number of non-overlapping occurrences of `term` in `docs[i]`.
3. Scan each document in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TermCounts([]string{"banana", "abc"}, "a")
Output: [3 1]
```

**Example 2:**

```
Input:  TermCounts([]string{"go go"}, "go")
Output: [2]
```

**Example 3:**

```
Input:  TermCounts(nil, "a")
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Capture what does not change** | `term` is fixed for the whole call, so capturing it is safe — unlike the loop variable `i`. |

## Hint

`i` and `doc` change every iteration, so pass them in. `term` is decided before
the loop, so the closure may read it directly.

## Validate

```bash
make verify
```

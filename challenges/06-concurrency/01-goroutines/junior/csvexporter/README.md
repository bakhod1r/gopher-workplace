# CSV Exporter

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A reporting job exports a query result to a delimited file. Each record is
rendered into its own line, records do not depend on one another, and the file
must keep the query's row order.

## Task

Implement `RenderRows` in [csvexporter.go](csvexporter.go) so that:

1. Return a slice with one rendered line per record, in record order.
2. Line `i` is the fields of `rows[i]` joined with `sep`; an empty record renders as the empty string.
3. Render each record in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RenderRows([][]string{{"a", "b"}}, ",")
Output: [a,b]
```

**Example 2:**

```
Input:  RenderRows([][]string{{}}, ",")
Output: []
```

**Example 3:**

```
Input:  RenderRows(nil, ",")
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Slice of slices** | Each inner slice is read by exactly one goroutine, so no two goroutines touch the same memory. |

## Hint

Pass `i` and `row` into the goroutine; `sep` is fixed for the whole export, so
capturing it from the enclosing scope is fine.

## Validate

```bash
make verify
```

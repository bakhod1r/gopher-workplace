# Partial Document Indexing

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

A reindex job pushes a batch of documents into the search cluster. Some documents are rejected — a missing body, a field the mapping does not accept — and the job's report has to say precisely which ones landed and which will be replayed. The rejects are a small minority, so the batch is never aborted over them.

## Task

Implement the exported function(s) in [docindexer.go](docindexer.go) so that:

1. Index each document in its own goroutine, joined with a `sync.WaitGroup`.
2. Attempt every document, whatever the others return.
3. Return the IDs of successfully indexed documents, sorted.
4. Return the IDs of rejected documents, sorted.
5. Both return values are empty non-nil slices when there is nothing to report.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  IndexDocuments([]Doc{{"a", "x"}, {"b", "y"}}, index)
Output: [a b], []
```

**Example 2:**

```
Input:  IndexDocuments([]Doc{{"zed", "x"}, {"beta", ""}}, index)
Output: [zed], [beta]
```

**Example 3:**

```
Input:  IndexDocuments(nil, index)
Output: [], []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Per-index flags** | A `[]bool` of outcomes is the cheapest race-free hand-off from workers to the fold. |
| 2 | **Partitioning after the join** | Splitting into two lists sequentially keeps both deterministic. |
| 3 | **Named return values** | `(indexed, failed []string)` documents which list is which at the call site. |
| 4 | **Sorted output** | Two lists that are compared by a job report must not shuffle between runs. |

## Hint

Do not build the two lists inside the goroutines. Record a boolean per document, then partition and sort in the parent after `wg.Wait()`.

## Validate

```bash
make verify
```

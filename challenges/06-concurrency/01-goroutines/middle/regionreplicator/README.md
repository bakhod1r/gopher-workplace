# Region Replication Tree

**Level:** middle
**Topic:** 06-concurrency → 01-goroutines

## Context

A write is replicated to every availability zone of every region. The natural shape is a tree: the coordinator owns one goroutine per region, and each of those owns one goroutine per zone inside it. Ownership matters — a region goroutine must not report itself finished while its zones are still writing.

## Task

Implement the exported function(s) in [regionreplicator.go](regionreplicator.go) so that:

1. Launch one goroutine per region, joined by an outer `sync.WaitGroup`.
2. Inside each region goroutine, launch one goroutine per zone joined by that region's own inner `sync.WaitGroup`.
3. A region goroutine returns only after its inner `Wait` returns.
4. Collect `"region/zone"` for every failing replication under a shared mutex.
5. Return the collected failures sorted, as an empty non-nil slice when there are none.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  ReplicateAll([]Region{{"eu", []string{"a", "bad-b"}}}, replicate)
Output: [eu/bad-b]
```

**Example 2:**

```
Input:  ReplicateAll([]Region{{"us", []string{"bad-1"}}, {"ap", []string{"bad-2"}}}, replicate)
Output: [ap/bad-2 us/bad-1]
```

**Example 3:**

```
Input:  ReplicateAll(nil, replicate)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Nested goroutine trees** | Each level owns its own WaitGroup; a parent joins its children before reporting done. |
| 2 | **Ownership and lifetime** | The outer `Wait` transitively covers every zone goroutine, so nothing outlives the call. |
| 3 | **Shared aggregation across levels** | One mutex protects the failure list no matter which level appends to it. |
| 4 | **Sorting for a stable runbook** | Completion order across two levels is doubly non-deterministic; the sort erases it. |

## Hint

Declare the inner `sync.WaitGroup` *inside* the region goroutine — a WaitGroup shared across regions would let one region wait on another's zones.

## Validate

```bash
make verify
```

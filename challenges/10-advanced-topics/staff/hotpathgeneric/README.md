# Type Parameters Instead Of Interface Boxing

**Level:** staff
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A metrics aggregator takes `[]any` so it can serve every integer width. Boxing the values costs one allocation per element, and the aggregator is called once per scrape per series.

## Task

Implement [hotpathgeneric.go](hotpathgeneric.go):

1. Sum `vals` and return the total as an int64.
2. Accept `int`, `int32`, `int64` and any named type based on them.
3. Zero allocations — nothing may be boxed.

Replace the stub body in [hotpathgeneric.go](hotpathgeneric.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Total([]int{1,2,3})
Output: 6
```

**Example 2:**

```
Input:  Total([]int64{1<<40, 1<<40})
Output: 2199023255552
```

_Explanation:_ The accumulator must be wide enough.

**Example 3:**

```
Input:  Total([]myInt{2,3})
Output: 5
```

_Explanation:_ `~int` admits named types.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Type parameters vs interfaces** | A generic call is compiled against the concrete type; an interface call boxes. |
| 2 | **Approximation constraints** | `~int` covers every type whose underlying type is int. |
| 3 | **Accumulator width** | Summing int32 into int32 overflows; int64 does not. |
| 4 | **Escape analysis with generics** | No boxing means no escape, so the loop stays in registers. |

## Hint

The constraint is already written. The body is the obvious loop — the point is what the signature bought you.

## Validate

```bash
make verify
```

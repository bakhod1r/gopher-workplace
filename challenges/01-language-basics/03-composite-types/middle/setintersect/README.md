# Set Intersection

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Finding common elements (shared permissions, mutual friends).

## Task

Implement `Intersect(a, b)` — sorted, unique, present in both.

## Examples

```go
Intersect([]int{1,2,3,4}, []int{2,4,6}) // => [2 4]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Set from a** | Membership test for b. |
| 2 | **Dedup result** | A value in b twice counts once. |
| 3 | **Sort output** | Deterministic order. |

## Hint

Set of a; for each b, if in a-set and not yet emitted, add; then sort.

## Validate

```bash
make verify
```

# Set Difference

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

"In A but not B" — revoked items, missing keys.

## Task

Implement `Diff(a, b)` — sorted, unique, in a not b.

## Examples

```go
Diff([]int{1,2,3,3,4}, []int{2,4}) // => [1 3]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Exclusion set** | Membership test against b. |
| 2 | **Dedup** | Emit each survivor once. |
| 3 | **Sort** | Deterministic order. |

## Hint

Set of b; for each a not in b and not already emitted, add; then sort.

## Validate

```bash
make verify
```

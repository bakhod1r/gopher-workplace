# Set Union

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Combining two tag sets into one sorted, unique list.

## Task

Implement `Union(a, b)` — sorted, de-duplicated.

## Examples

```go
Union([]int{3,1,2}, []int{2,4,1}) // => [1 2 3 4]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Set membership** | `map[int]struct{}`. |
| 2 | **Collect + sort** | Keys to slice, then `sort.Ints`. |
| 3 | **Dedup** | Set removes duplicates automatically. |

## Hint

Insert all of a and b into a set, collect keys, `sort.Ints`.

## Validate

```bash
make verify
```

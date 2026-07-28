# Top K Frequent

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

"Top N" analytics: count occurrences, then rank.

## Task

Implement `TopK(xs, k)` — k most frequent, ties broken alphabetically.

## Examples

```go
TopK([]string{"a","b","a","c","b","a"}, 2) // => [a b]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Count map** | `m[x]++`. |
| 2 | **Custom sort** | `sort.Slice` with count then name. |
| 3 | **Clamp k** | Not more than distinct count. |

## Hint

Count into a map, collect keys, `sort.Slice` by `(count desc, key asc)`, take
first k.

## Validate

```bash
make verify
```

# Group By Key

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A key function plus a map builds buckets; appending to `m[k]` works even when
the key is new (nil slice appends fine).

## Task

Implement `GroupBy` in [groupby.go](groupby.go).

Do **not** change the function signature or the tests.

## Examples

```go
GroupBy([]int{1,2,3,4}, mod2) // => {0:[2 4], 1:[1 3]}
GroupBy(nil, k)               // => map[]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Key function** | `key(x)` chooses the bucket. |
| 2 | **Map of slices** | `m[k] = append(m[k], x)`. |
| 3 | **nil-slice append** | Appending to a missing key's nil slice is valid. |

## Hint

Init `m := map[int][]int{}`; range `xs` doing `k := key(x); m[k] = append(m[k], x)`.

## Validate

```bash
make verify
```

# Set Membership Value

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

The set uses `map[int]bool`, but members are stored as `false`. Membership is
tested with `inB[x]` (the value), which is always false — so the intersection is
empty.

## Task

Fix the store between the markers in
[emptystructset.go](emptystructset.go) so membership is true.

## Examples

```go
Intersect([]int{1,2,3}, []int{2,3,5}) // => [2 3]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bool set** | Store `true` to mark membership. |
| 2 | **Value vs presence** | `inB[x]` reads the stored bool. |
| 3 | **Alternative** | `map[int]struct{}` + comma-ok. |

## Hint

`inB[x] = true`.

## Validate

```bash
make verify
```

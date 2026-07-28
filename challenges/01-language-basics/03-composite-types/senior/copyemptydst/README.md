# copy Into a Length-0 Slice

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`copy` copies `min(len(dst), len(src))` elements. The destination is made with
**length 0** (only capacity), so `copy` copies nothing and `Clone` returns empty.

## Task

Fix the `make` between the markers in [copyemptydst.go](copyemptydst.go) to give
the destination the right length.

## Examples

```go
Clone([]int{1,2,3}) // => [1 2 3]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **copy semantics** | Copies `min(len(dst),len(src))`. |
| 2 | **len vs cap** | Capacity alone doesn't hold elements. |
| 3 | **Size dst** | `make([]int, len(xs))`. |

## Hint

`make([]int, len(xs))` (length, not just capacity).

## Validate

```bash
make verify
```

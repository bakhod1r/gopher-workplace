# Length vs Capacity for Indexing

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

`make([]int, 0, len(xs))` has capacity but **length 0**, so `out[i] = ...` panics
(index out of range). Either index into a full-length slice, or `append`.

## Task

Fix the build between the markers in
[preallocindex.go](preallocindex.go).

## Examples

```go
Doubled([]int{1,2,3}) // => [2 4 6]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **len vs cap** | You can only index within length. |
| 2 | **Two idioms** | `make(_, len)`+index, or `make(_,0,cap)`+append. |
| 3 | **Panic** | Indexing past length panics. |

## Hint

Either `out := make([]int, len(xs))` and index, or keep cap and `append`.

## Validate

```bash
make verify
```

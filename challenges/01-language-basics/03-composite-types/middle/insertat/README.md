# Insert at Index

**Level:** middle
**Topic:** 01-language-basics → 03-composite-types

## Context

Inserting into an ordered slice: split at the index and stitch the value in.

## Task

Implement `InsertAt(xs, i, v)`, clamping `i` to `[0, len]`.

## Examples

```go
InsertAt([]int{1,2,3}, 1, 9) // => [1 9 2 3]
InsertAt([]int{1,2,3}, 10,9) // => [1 2 3 9]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Split & stitch** | `append(head, v)` then tail. |
| 2 | **Aliasing hazard** | `append(xs[:i], ...)` can clobber the tail. |
| 3 | **Clamp index** | i in `[0, len]`. |

## Hint

Safest: `out := append([]int{}, xs[:i]...); out = append(out, v); out = append(out, xs[i:]...)`.

## Validate

```bash
make verify
```

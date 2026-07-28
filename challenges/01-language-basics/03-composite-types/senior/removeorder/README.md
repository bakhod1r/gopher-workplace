# Delete Off-By-One

**Level:** senior
**Topic:** 01-language-basics → 03-composite-types

## Context

The delete idiom joins the parts before and after index `i`. The tail slice
starts at `i` instead of `i+1`, so the element is copied back, not removed.

## Task

Fix the tail expression between the markers in
[removeorder.go](removeorder.go).

## Examples

```go
RemoveAt([]int{1,2,3,4}, 1) // => [1 3 4]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two-half join** | Head `[:i]`, tail `[i+1:]`. |
| 2 | **Skip the target** | Tail starts *after* `i`. |
| 3 | **Spread** | `...` expands the tail. |

## Hint

`append(xs[:i], xs[i+1:]...)`.

## Validate

```bash
make verify
```

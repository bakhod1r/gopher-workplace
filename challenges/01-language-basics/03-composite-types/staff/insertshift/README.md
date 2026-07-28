# Insert Shift Direction

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

To insert at `i`, the elements from `i` onward must shift **right** by one. The
code does `copy(xs[i:], xs[i+1:])`, which shifts left, dropping the tail.

## Task

Fix the copy between the markers in [insertshift.go](insertshift.go).

## Examples

```go
InsertAt([]int{1,2,3,4}, 1, 9) // => [1 9 2 3 4]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Shift right** | `copy(xs[i+1:], xs[i:])`. |
| 2 | **Overlap-safe** | `copy` handles the overlap. |
| 3 | **Grow then shift** | Append a slot, then open the gap. |

## Hint

`copy(xs[i+1:], xs[i:])`.

## Validate

```bash
make verify
```

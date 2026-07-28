# copy Direction for Shift

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Shifting left drops `xs[0]`: element `i+1` moves to `i`. The code does
`copy(xs[1:], xs)`, which shifts **right** (overlapping copy toward higher
indices), duplicating the first element.

## Task

Fix the copy between the markers in
[copyshiftleft.go](copyshiftleft.go).

## Examples

```go
ShiftLeft([]int{1,2,3,4}) // => [2 3 4 0]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **copy on overlap** | `copy` handles overlap correctly. |
| 2 | **Source/dest** | Left shift: `copy(xs, xs[1:])`. |
| 3 | **Fill vacated** | Zero the last slot. |

## Hint

`copy(xs, xs[1:])`.

## Validate

```bash
make verify
```

# Array Key Coordinate Order

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Arrays are comparable, so `[2]int` is a valid map key. But the key is built with
the coordinates **swapped** (`{c[1], c[0]}`), so `(1,2)` and `(2,1)` are conflated.

## Task

Fix the key between the markers in [maparraykey.go](maparraykey.go).

## Examples

```go
CountCells([{1,2},{1,2},{2,1}]) // {1,2}:2, {2,1}:1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Arrays are comparable** | Usable as map keys. |
| 2 | **Slices are not** | `[]int` can't be a key. |
| 3 | **Key identity** | Field order defines the key. |

## Hint

`m[[2]int{c[0], c[1]}]++`.

## Validate

```bash
make verify
```

# Struct Map Key Fields

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

A struct of comparable fields is a valid map key. But the key is built with `X`
and `Y` swapped (`Point{p.Y, p.X}`), so `(1,2)` and `(2,1)` collide.

## Task

Fix the key between the markers in [keyfromstruct.go](keyfromstruct.go).

## Examples

```go
Count([{1,2},{1,2},{2,1}]) // {1,2}:2, {2,1}:1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Structs as keys** | Comparable structs are valid keys. |
| 2 | **Field identity** | Field values define the key. |
| 3 | **No slice fields** | A slice field makes it non-comparable. |

## Hint

`m[Point{p.X, p.Y}]++` (or `m[p]++`).

## Validate

```bash
make verify
```

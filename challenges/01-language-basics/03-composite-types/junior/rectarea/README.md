# Rectangle Struct

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

A geometry type with methods: area, and a scale that returns a new value without
mutating the receiver.

## Task

Implement `Area()` and `Scale(factor)` on `Rect`.

## Examples

```go
Rect{3,4}.Area()   // => 12
Rect{2,3}.Scale(2) // => {4 6}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Value receiver** | `func (r Rect)` gets a copy. |
| 2 | **Return new struct** | Scale builds a fresh Rect. |
| 3 | **Struct equality** | `==` compares fields. |

## Hint

`Area`: `r.W * r.H`. `Scale`: `Rect{r.W*factor, r.H*factor}`.

## Validate

```bash
make verify
```

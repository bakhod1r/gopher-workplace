# Value Receiver Returns Copy

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A value receiver is a copy; returning a new value (rather than mutating) is the
idiomatic immutable style for small structs like points.

## Task

Implement the `Scaled` method in [vscale.go](vscale.go).

Do **not** change the function signature or the tests.

## Examples

```go
Point{1,2}.Scaled(3) // => {3, 6}
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Value receiver** | The method gets a copy. |
| 2 | **Return new value** | Build and return a fresh Point. |
| 3 | **Immutability** | The original is untouched. |

## Hint

`return Point{X: p.X * k, Y: p.Y * k}`.

## Validate

```bash
make verify
```

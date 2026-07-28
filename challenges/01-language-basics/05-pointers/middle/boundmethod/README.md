# Bind a Pointer Method

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A method value `c.Inc` bound on a pointer receiver captures the pointer, so the
returned function keeps mutating the same counter.

## Task

Implement `Bind` in [boundmethod.go](boundmethod.go).

Do **not** change the function signature or the tests.

## Examples

```go
inc := Bind(c); inc(); inc() // c.N = 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Method value** | `c.Inc` is a callable bound to c. |
| 2 | **Pointer capture** | The bound receiver is the pointer. |
| 3 | **Shared state** | All calls hit the same counter. |

## Hint

`return c.Inc` (the method value bound to the pointer).

## Validate

```bash
make verify
```

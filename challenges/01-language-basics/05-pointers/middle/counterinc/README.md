# Pointer Receiver Increment

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A method with a pointer receiver `(c *Counter)` can mutate the receiver; a value
receiver would modify a copy and lose the change.

## Task

Implement the `Inc` method in [counterinc.go](counterinc.go).

Do **not** change the function signature or the tests.

## Examples

```go
c := &Counter{}; c.Inc() // c.N = 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer receiver** | `(c *Counter)` mutates the original. |
| 2 | **Field update** | `c.N++`. |
| 3 | **Value vs pointer receiver** | Value receiver can't persist changes. |

## Hint

`c.N++`.

## Validate

```bash
make verify
```

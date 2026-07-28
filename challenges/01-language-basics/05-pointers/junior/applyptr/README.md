# Apply Function In Place

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Combining a pointer with a function value transforms the caller's variable in
place with caller-chosen logic.

## Task

Implement `Apply` in [applyptr.go](applyptr.go).

Do **not** change the function signature or the tests.

## Examples

```go
Apply(&x, square) // x = square(x)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Read, transform, write** | `*p = f(*p)`. |
| 2 | **Function parameter** | Caller supplies the operation. |
| 3 | **In-place** | Result stored back. |

## Hint

`*p = f(*p)`.

## Validate

```bash
make verify
```

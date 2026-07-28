# Apply Discards the Result

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

`f(*p)` computes the new value but throws it away. To update the caller's
variable you must store the result back: `*p = f(*p)`.

## Task

Fix [applyinplace.go](applyinplace.go) so the transform is stored.

Do **not** change the function signature or the tests.

## Examples

```go
Apply(&x, square) // *p = square(*p)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Store the result** | `*p = f(*p)`. |
| 2 | **Discarded computation** | `f(*p)` alone changes nothing. |
| 3 | **Write through the pointer** | The pointee must be reassigned. |

## Hint

Store it back: `*p = f(*p)`.

## Validate

```bash
make verify
```

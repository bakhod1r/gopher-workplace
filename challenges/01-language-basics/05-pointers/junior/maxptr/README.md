# Return Pointer to Larger

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Returning a pointer lets the caller keep mutating the chosen variable. Compare
the pointed-to values, return the address.

## Task

Implement `MaxPtr` in [maxptr.go](maxptr.go).

Do **not** change the function signature or the tests.

## Examples

```go
MaxPtr(&3, &8) // => &8
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Compare pointees** | `*a` vs `*b`. |
| 2 | **Return an address** | Result type `*int`. |
| 3 | **Tie rule** | Return a when equal. |

## Hint

`if *b > *a { return b }; return a`.

## Validate

```bash
make verify
```

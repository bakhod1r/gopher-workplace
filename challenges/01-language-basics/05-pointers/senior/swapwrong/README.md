# Swapping Pointers, Not Values

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

`a, b = b, a` swaps the local pointer parameters, which are copies; the caller's
variables are untouched. Swap the POINTED-TO values with `*a, *b = *b, *a`.

## Task

Fix [swapwrong.go](swapwrong.go).

Do **not** change the function signature or the tests.

## Examples

```go
x,y := 1,2; Swap(&x,&y) // x=2, y=1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer params are copies** | Reassigning them is local. |
| 2 | **Dereference to swap values** | `*a, *b = *b, *a`. |
| 3 | **Caller visibility** | Only value writes reach the caller. |

## Hint

Swap the values, not the pointers: `*a, *b = *b, *a`.

## Validate

```bash
make verify
```

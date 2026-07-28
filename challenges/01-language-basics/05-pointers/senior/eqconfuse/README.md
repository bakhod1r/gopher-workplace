# Identity vs Value Confusion

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

`*a == *b` compares the VALUES, which can be equal for different variables.
Identity is `a == b` — comparing the addresses.

## Task

Fix [eqconfuse.go](eqconfuse.go) to compare identity.

Do **not** change the function signature or the tests.

## Examples

```go
Same(&x, &x) // => true
Same(&x, &y) // => false even if *x == *y
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Identity vs value** | `a == b` (address) vs `*a == *b` (value). |
| 2 | **Distinct storage** | Equal values, different addresses. |
| 3 | **Pointer comparison** | Compare the pointers directly. |

## Hint

Compare addresses: `return a == b`.

## Validate

```bash
make verify
```

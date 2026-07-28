# Increment Through Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

A pointer parameter `*int` refers to the caller's variable. Writing through it
with `*p` changes the original.

## Task

Implement `Inc` in [incptr.go](incptr.go).

Do **not** change the function signature or the tests.

## Examples

```go
x := 41; Inc(&x) // x becomes 42
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer parameter** | `*int` aliases the caller's variable. |
| 2 | **Dereference to write** | `*p` reads/writes the pointee. |
| 3 | **Address-of** | The caller passes `&x`. |

## Hint

`*p++` (or `*p = *p + 1`).

## Validate

```bash
make verify
```

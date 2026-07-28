# Toggle a Bool Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-basics_

## Context

Pointers work for any type, including bool. Flip with `!`.

## Task

Implement `Toggle` in [toggle.go](toggle.go).

Do **not** change the function signature or the tests.

## Examples

```go
b := false; Toggle(&b) // b=true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bool pointer** | `*bool` aliases a flag. |
| 2 | **Logical not** | `*p = !*p`. |
| 3 | **Toggle** | Two flips restore. |

## Hint

`*p = !*p`.

## Validate

```bash
make verify
```

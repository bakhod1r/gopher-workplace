# Mutate Struct Field via Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Through a struct pointer you can update a field in place. Go lets you write
`c.Count` on a pointer — it auto-dereferences.

## Task

Implement `Grow` in [bumpfield.go](bumpfield.go).

Do **not** change the function signature or the tests.

## Examples

```go
c := Cart{Count:2}; Grow(&c) // c.Count = 3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Struct pointer** | `*Cart` aliases the caller's struct. |
| 2 | **Auto-deref field access** | `c.Count` works on a pointer. |
| 3 | **In-place field update** | `c.Count++`. |

## Hint

`c.Count++` (Go auto-dereferences the pointer for field access).

## Validate

```bash
make verify
```

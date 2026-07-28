# Append via Slice Pointer

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Appending to a local copy `s` and discarding it leaves the caller's slice
unchanged. Assign the append result back through the pointer: `*sp = append(*sp, vs...)`.

## Task

Fix [dblappend.go](dblappend.go) so the caller's slice grows.

Do **not** change the function signature or the tests.

## Examples

```go
var xs []int; Extend(&xs, 1,2,3) // xs = [1 2 3]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Assign through the pointer** | `*sp = append(*sp, vs...)`. |
| 2 | **Local copy lost** | Appending to `s` doesn't update `*sp`. |
| 3 | **Slice header propagation** | The pointer carries the new header. |

## Hint

Write back through the pointer: `*sp = append(*sp, vs...)`.

## Validate

```bash
make verify
```

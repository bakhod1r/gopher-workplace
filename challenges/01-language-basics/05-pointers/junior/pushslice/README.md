# Append Through Slice Pointer

**Level:** junior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

append may return a new slice header. To have the caller see growth, pass a
pointer to the slice and assign the append result through it.

## Task

Implement `Push` in [pushslice.go](pushslice.go).

Do **not** change the function signature or the tests.

## Examples

```go
var xs []int; Push(&xs, 1) // xs = [1]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer to slice** | `*[]int` aliases the caller's slice header. |
| 2 | **Reassign through pointer** | `*sp = append(*sp, v)`. |
| 3 | **Header update** | The caller sees the new len/cap. |

## Hint

`*sp = append(*sp, v)`.

## Validate

```bash
make verify
```

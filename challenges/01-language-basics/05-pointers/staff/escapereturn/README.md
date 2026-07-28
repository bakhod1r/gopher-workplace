# Return Pointer to the Element

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _memory-management_

## Context

`&v` addresses a local copy of the element (which escapes to the heap but is
detached from the slice). To return an aliasing pointer, take the element's
address directly: `&xs[best]`.

## Task

Fix [escapereturn.go](escapereturn.go) so the returned pointer aliases the slice.

Do **not** change the function signature or the tests.

## Examples

```go
p := MaxPtr(xs); *p = 0 // xs's max element becomes 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Address of a copy** | `&v` detaches from the slice. |
| 2 | **Address of the element** | `&xs[best]` aliases the backing array. |
| 3 | **Escape analysis** | Both escape, but only one aliases xs. |

## Hint

Return the element address: `return &xs[best]`.

## Validate

```bash
make verify
```

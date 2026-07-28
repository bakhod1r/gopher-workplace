# Push Discards Append Result

**Level:** senior
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

`append` returns a possibly-new slice header; discarding it leaves `s.data`
unchanged. Assign the result back to the field.

## Task

Fix [pushnoassign.go](pushnoassign.go).

Do **not** change the function signature or the tests.

## Examples

```go
s.Push(1); s.Push(2) // len == 2
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **append returns a header** | The field must be reassigned. |
| 2 | **Reassign the field** | `s.data = append(s.data, v)`. |
| 3 | **Pointer receiver** | Mutation persists on the struct. |

## Hint

Assign back: `s.data = append(s.data, v)`.

## Validate

```bash
make verify
```

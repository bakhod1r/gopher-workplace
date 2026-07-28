# List to Slice

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

Flattening a linked list into a slice is a straight traversal, appending each
value.

## Task

Implement `ToSlice` in [listslice.go](listslice.go).

Do **not** change the function signature or the tests.

## Examples

```go
ToSlice(1->2->3) // => [1 2 3]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Traverse** | Walk Next until nil. |
| 2 | **Append values** | Collect Val. |
| 3 | **Empty case** | nil -> empty slice. |

## Hint

`for n := head; n != nil; n = n.Next { out = append(out, n.Val) }`.

## Validate

```bash
make verify
```

# Stack Push/Pop

**Level:** middle
**Topic:** 01-language-basics → 05-pointers · _pointers-with-structs_

## Context

A stack Push mutates the receiver's slice, so it needs a pointer receiver and
must assign the append result back.

## Task

Implement the `Push` method in [ptrstack.go](ptrstack.go).

Do **not** change the function signature or the tests.

## Examples

```go
s.Push(1); s.Push(2) // stack: [1 2]
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pointer receiver** | Mutation must persist. |
| 2 | **Append to field** | `s.data = append(s.data, v)`. |
| 3 | **Encapsulated state** | The slice lives in the struct. |

## Hint

`s.data = append(s.data, v)`.

## Validate

```bash
make verify
```

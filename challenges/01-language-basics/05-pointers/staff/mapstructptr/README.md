# Mutate Pointee From Map Struct

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _with-maps-and-slices_

## Context

Even though the `Ref` value is a copy from the map, its `P` field still points
at the shared int. Mutate through the pointer: `*r.P++`. The bug does a no-op
self-assignment.

## Task

Fix [mapstructptr.go](mapstructptr.go) to increment the shared int.

Do **not** change the function signature or the tests.

## Examples

```go
BumpVia(m, 1) // *m[1].P += 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Copied struct, shared pointee** | The copy's pointer still aliases. |
| 2 | **Mutate through the pointer** | `*r.P++`. |
| 3 | **Map value addressability** | You can't take &m[k], but the pointer inside works. |

## Hint

Increment through the pointer: `*r.P++`.

## Validate

```bash
make verify
```

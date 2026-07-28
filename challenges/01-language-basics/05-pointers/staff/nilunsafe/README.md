# Nil Check Before Unsafe Read

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _nil-pointer-dereference_

## Context

Even with unsafe pointers, a nil check must PRECEDE the dereference. Reading
`*(*int)(p)` first panics when p is nil.

## Task

Fix [nilunsafe.go](nilunsafe.go) so a nil pointer returns the default without panicking.

Do **not** change the function signature or the tests.

## Examples

```go
ReadOr(nil, 7) // => 7
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Guard before deref** | Check nil first, even for unsafe. |
| 2 | **Order of operations** | The panic is at the read. |
| 3 | **Default on nil** | Return def. |

## Hint

Check nil first: `if p == nil { return def }; return *(*int)(p)`.

## Validate

```bash
make verify
```

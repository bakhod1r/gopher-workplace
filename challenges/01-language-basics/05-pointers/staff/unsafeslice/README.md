# Build a Slice From a Pointer

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

`unsafe.Slice(ptr, n)` takes n as an ELEMENT count. `unsafe.Sizeof(*p)` is 16
bytes, not 4 elements; pass the element count (4, or `len(p)`).

## Task

Fix [unsafeslice.go](unsafeslice.go) to use the element count.

Do **not** change the function signature or the tests.

## Examples

```go
View(&[4]int32{...}) // len 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Element count** | unsafe.Slice's second arg is a count. |
| 2 | **Bytes vs elements** | Sizeof gives bytes. |
| 3 | **Correct length** | Use `len(p)` or the element count. |

## Hint

Pass the element count: `unsafe.Slice(&p[0], len(p))`.

## Validate

```bash
make verify
```

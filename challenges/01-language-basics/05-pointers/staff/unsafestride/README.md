# Pointer Arithmetic Stride

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

`unsafe.Add(ptr, n)` advances by n BYTES. To step i elements you must multiply
by the element size: `i * unsafe.Sizeof(arr[0])` (4 for int32).

## Task

Fix the offset in [unsafestride.go](unsafestride.go).

Do **not** change the function signature or the tests.

## Examples

```go
At(&[10 20 30 40], 2) // => 30
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Byte-based offset** | unsafe.Add moves by bytes. |
| 2 | **Element stride** | Multiply by the element size. |
| 3 | **unsafe.Sizeof** | Gives the element width. |

## Hint

Scale by the element size: `unsafe.Add(base, uintptr(i)*unsafe.Sizeof(arr[0]))`.

## Validate

```bash
make verify
```

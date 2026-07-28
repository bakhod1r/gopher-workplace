# First Byte of a String

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

`&s` addresses the string HEADER (data ptr + len), not the characters. The data
pointer is `unsafe.StringData(s)` (equivalently the first byte's address).

## Task

Fix [stringbyte.go](stringbyte.go) to point at the string data.

Do **not** change the function signature or the tests.

## Examples

```go
FirstByte("Xyz") // => 'X'
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **String header vs data** | `&s` addresses the header. |
| 2 | **unsafe.StringData** | Returns the data pointer. |
| 3 | **Read-only** | String bytes must not be mutated. |

## Hint

Use the data pointer: `p := unsafe.Pointer(unsafe.StringData(s))`.

## Validate

```bash
make verify
```

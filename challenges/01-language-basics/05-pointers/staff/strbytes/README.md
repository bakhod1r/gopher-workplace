# String Length From Header

**Level:** staff
**Topic:** 01-language-basics → 05-pointers · _unsafe-pointer_

## Context

A string header is (Data pointer, Len). The bug reads the Data word (an
address); the length is the second word, `h.Len`.

## Task

Fix [strbytes.go](strbytes.go) to read the length field.

Do **not** change the function signature or the tests.

## Examples

```go
ByteLen("hello") // => 5
ByteLen("héllo")  // => 6
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **String header layout** | (data pointer, length). |
| 2 | **Length word** | The second field is the length. |
| 3 | **Data vs length** | Data is an address, not the count. |

## Hint

Read the length word: `return h.Len`.

## Validate

```bash
make verify
```

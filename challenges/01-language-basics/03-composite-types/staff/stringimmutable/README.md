# Strings Are Immutable

**Level:** staff
**Topic:** 01-language-basics → 03-composite-types

## Context

Strings are immutable, so the code copies to a `[]byte`, mutates that, then
mistakenly returns the **original** `s` — the edits are lost.

## Task

Fix the return between the markers in
[stringimmutable.go](stringimmutable.go) to return the modified bytes.

## Examples

```go
Upper("hello") // => "HELLO"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **String immutability** | Bytes of a string can't change. |
| 2 | **Mutate via []byte** | Edit a copy, then convert back. |
| 3 | **Convert back** | `string(b)` is the result. |

## Hint

`return string(b)`.

## Validate

```bash
make verify
```

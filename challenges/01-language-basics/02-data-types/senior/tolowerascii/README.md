# ASCII Lowercase Range

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A header normalizer lowercases ASCII, but the range check only has a lower bound
(`c >= 'A'`), so `'['`, `'\'`, `']'` — the bytes right after `'Z'` — get shifted
too, corrupting the text.

## Task

Fix the condition between the markers in [tolowerascii.go](tolowerascii.go) to
bound `A..Z`.

## Examples

```go
Lower("Hello") // => "hello"
Lower("a[b]")  // => "a[b]"  (brackets unchanged)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **ASCII ranges** | `'A'..'Z'` is 65..90; `'['` is 91. |
| 2 | **Both bounds** | Need `>= 'A' && <= 'Z'`. |
| 3 | **Case offset** | Lowercase is +32. |

## Hint

`if c >= 'A' && c <= 'Z'`.

## Validate

```bash
make verify
```

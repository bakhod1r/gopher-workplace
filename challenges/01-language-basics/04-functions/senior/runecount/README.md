# Range Over String Yields Runes

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`len(s)` on a string is its byte length; multibyte UTF-8 runes count as 2–4
bytes. Ranging a string (or `utf8.RuneCountInString`) counts actual code points.

## Task

Fix [runecount.go](runecount.go) to count runes, not bytes.

Do **not** change the function signature or the tests.

## Examples

```go
CharCount("héllo") // => 5
CharCount("日本語")   // => 3
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bytes vs runes** | `len(s)` is bytes; runes may be multibyte. |
| 2 | **Range over string** | `for range s` steps rune by rune. |
| 3 | **utf8 helpers** | `utf8.RuneCountInString` is the direct call. |

## Hint

Count with a `for range s` loop, or `return utf8.RuneCountInString(s)`.

## Validate

```bash
make verify
```

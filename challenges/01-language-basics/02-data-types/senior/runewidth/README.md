# First Rune Byte Width

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A streaming UTF-8 decoder needs to advance by the width of the current rune, but
the function returns `len(s)` — the whole remaining length — so the decoder
jumps to the end instead of the next rune.

## Task

Fix the return between the markers in [runewidth.go](runewidth.go) to return the
first rune's byte size.

## Examples

```go
FirstWidth("é")   // => 2
FirstWidth("日本") // => 3  (first rune only)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **utf8.DecodeRuneInString** | Returns the rune and its byte size. |
| 2 | **Rune width** | 1–4 bytes per rune. |
| 3 | **Advance step** | Decoders step by `size`. |

## Hint

`return size`.

## Validate

```bash
make verify
```

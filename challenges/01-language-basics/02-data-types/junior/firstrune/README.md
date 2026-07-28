# First Rune

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Indexing a string with `s[0]` gives a **byte**, which splits multi-byte
characters. Ranging over a string yields **runes** (Unicode code points).

## Task

Implement `First(s)` returning the first rune. Empty string returns rune `0`.

## Examples

```go
First("hello") // => 'h'
First("étage") // => 'é'  (not the first byte)
First("日本")   // => '日'
First("")      // => 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bytes vs runes** | `s[0]` is a byte; a rune may span several bytes. |
| 2 | **for range on string** | Yields `(byteIndex, rune)` decoding UTF-8. |
| 3 | **Early return** | Return on the first iteration. |

## Hint

`for _, r := range s { return r }` then `return 0`.

## Validate

```bash
make verify
```

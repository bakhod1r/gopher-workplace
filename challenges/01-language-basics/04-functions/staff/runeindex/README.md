# Index Into Runes, Not Bytes

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`s[i]` indexes BYTES; for multibyte UTF-8 the i-th byte is not the i-th
character. Convert to `[]rune` (or decode) to index by character position.

## Task

Fix [runeindex.go](runeindex.go) so it returns the i-th rune.

Do **not** change the function signature or the tests.

## Examples

```go
CharAt("héllo", 2) // => 'l'
CharAt("日本語", 1) // => '本'
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Byte vs rune indexing** | `s[i]` is a byte, not a code point. |
| 2 | **[]rune conversion** | `[]rune(s)[i]` indexes by character. |
| 3 | **UTF-8 width** | Runes span 1–4 bytes. |

## Hint

Convert first: `return []rune(s)[i]`.

## Validate

```bash
make verify
```

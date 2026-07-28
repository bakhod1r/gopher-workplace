# Rune Count

**Level:** junior
**Topic:** 01-language-basics → 02-data-types
**Estimated time:** 10 min

## Context

A form validator must limit a username to N characters. A first attempt used
`len(s)`, but users with accented or non-Latin names hit the limit too early —
`len` counts bytes, not characters.

## Task

Implement `Count` in [runecount.go](runecount.go) so it returns the number of
**runes** (characters) in `s`, correct for multi-byte UTF-8 input.

Do **not** change the function signature or the tests.

## Examples

```go
Count("abc")   // => 3
Count("héllo") // => 5   (é is 2 bytes, 1 rune)
Count("日本")   // => 2
Count("a🙂b")   // => 3
Count("")      // => 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Strings are bytes** | A Go string is a read-only sequence of bytes encoded as UTF-8. |
| 2 | **`len` is a byte count** | `len(s)` returns bytes, so a 2-byte `é` inflates it past the character count. |
| 3 | **Runes** | A `rune` is a Unicode code point. `[]rune(s)` decodes the string, and `range` over a string yields runes too. |

## Hint

`len(s)` counts bytes. Convert to `[]rune(s)` and take its length, or count with
a `range` loop over the string (which iterates rune by rune).

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

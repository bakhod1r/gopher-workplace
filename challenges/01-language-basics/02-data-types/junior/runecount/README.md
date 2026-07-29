# Rune Count

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

A form validator must limit a username to N characters. A first attempt used
`len(s)`, but users with accented or non-Latin names hit the limit too early —
`len` counts bytes, not characters.

## Task

Implement `Count` in [runecount.go](runecount.go) so it returns the number of
**runes** (characters) in `s`, correct for multi-byte UTF-8 input.

Do **not** change the function signature or the tests.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Count("abc")
Output: 3
```

_Explanation:_ ASCII, bytes==runes.

**Example 2:**

```
Input:  Count("hello" with accent)
Output: 5
```

_Explanation:_ The accented char is 2 bytes but 1 rune; len would give 6.

**Example 3:**

```
Input:  Count("CJK pair")
Output: 2
```

_Explanation:_ 3 bytes each, 2 runes.

**Example 4:**

```
Input:  Count("a<emoji>b")
Output: 3
```

_Explanation:_ The emoji is 4 bytes, still 1 rune.

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
make verify
```

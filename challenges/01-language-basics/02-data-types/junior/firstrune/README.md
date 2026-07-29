# First Rune

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Indexing a string with `s[0]` gives a **byte**, which splits multi-byte
characters. Ranging over a string yields **runes** (Unicode code points).

## Task

Implement `First(s)` returning the first rune. Empty string returns rune `0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  First("hello")
Output: 'h'
```

_Explanation:_ First rune.

**Example 2:**

```
Input:  First("etage" with accent)
Output: 'e-acute'
```

_Explanation:_ Ranging decodes the 2-byte accented char as one rune.

**Example 3:**

```
Input:  First("day" CJK)
Output: first CJK rune
```

_Explanation:_ 3-byte rune returned whole.

**Example 4:**

```
Input:  First("")
Output: 0
```

_Explanation:_ Empty string, loop never runs, return rune 0.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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

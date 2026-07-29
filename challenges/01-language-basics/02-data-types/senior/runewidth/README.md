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

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "a"
Output: 1
```

**Example 2:**

```
Input:  "é"
Output: 2
```

**Example 3:**

```
Input:  "日本"
Output: 3
```

_Explanation:_ first rune only

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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

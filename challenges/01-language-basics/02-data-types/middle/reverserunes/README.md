# Reverse by Runes

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Reversing bytes corrupts multi-byte UTF-8 characters. Convert to `[]rune` first,
reverse, convert back.

## Task

Implement `Reverse(s)` reversing runes, keeping characters intact.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Reverse("hello")
Output: "olleh"
```

_Explanation:_ ASCII reversal

**Example 2:**

```
Input:  Reverse("café")
Output: "éfac"
```

_Explanation:_ é kept intact (rune-aware)

**Example 3:**

```
Input:  Reverse("日本語")
Output: "語本日"
```

_Explanation:_ multi-byte kanji reversed as whole runes

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **[]rune(s)** | Decodes UTF-8 into code points. |
| 2 | **Two-pointer swap** | Swap ends moving inward. |
| 3 | **string([]rune)** | Re-encodes to UTF-8. |

## Hint

`r := []rune(s)`; swap `r[i], r[j]` inward; `return string(r)`.

## Validate

```bash
make verify
```

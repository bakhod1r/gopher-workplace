# Title Case

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Title-casing upper-cases each word's first letter and lower-cases the rest —
case conversion on ASCII letters via arithmetic or `unicode` helpers.

## Task

Implement `Title(s)` for space-separated words.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Title("hello world")
Output: "Hello World"
```

_Explanation:_ first letter of each word capitalized

**Example 2:**

```
Input:  Title("GO is FUN")
Output: "Go Is Fun"
```

_Explanation:_ rest of each word lowercased

**Example 3:**

```
Input:  Title("a")
Output: "A"
```

_Explanation:_ single letter uppercased

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Case conversion** | `unicode.ToUpper/ToLower` on runes. |
| 2 | **Word boundary** | Track "start of word" while scanning. |
| 3 | **Building output** | strings.Builder or []rune. |

## Hint

Track a `startOfWord` flag; upper the first letter, lower the others; reset on
space.

## Validate

```bash
make verify
```

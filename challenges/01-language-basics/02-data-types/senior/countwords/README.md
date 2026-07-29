# Robust Word Count

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

An analytics job counts words with `strings.Split(s, " ")`, which turns every
run of spaces into empty strings — so `"  a   b  "` counts as 6, not 2.

## Task

Fix the single line between the markers in [countwords.go](countwords.go) to
count only non-empty words.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "hello world"
Output: 2
```

**Example 2:**

```
Input:  "  a   b  "
Output: 2
```

_Explanation:_ runs of spaces are not words

**Example 3:**

```
Input:  "   "
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Split vs Fields** | `strings.Fields` splits on runs of whitespace. |
| 2 | **Empty tokens** | `Split` on a single space yields empties. |
| 3 | **Whitespace kinds** | `Fields` also handles tabs/newlines. |

## Hint

`return len(strings.Fields(s))`.

## Validate

```bash
make verify
```

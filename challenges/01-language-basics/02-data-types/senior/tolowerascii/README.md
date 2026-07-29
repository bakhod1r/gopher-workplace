# ASCII Lowercase Range

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A header normalizer lowercases ASCII, but the range check only has a lower bound
(`c >= 'A'`), so `'['`, `'\'`, `']'` — the bytes right after `'Z'` — get shifted
too, corrupting the text.

## Task

Fix the condition between the markers in [tolowerascii.go](tolowerascii.go) to
bound `A..Z`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "Hello"
Output: "hello"
```

**Example 2:**

```
Input:  "a[b]"
Output: "a[b]"
```

_Explanation:_ '[' just past 'Z' must stay

**Example 3:**

```
Input:  "MixED_case"
Output: "mixed_case"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **ASCII ranges** | `'A'..'Z'` is 65..90; `'['` is 91. |
| 2 | **Both bounds** | Need `>= 'A' && <= 'Z'`. |
| 3 | **Case offset** | Lowercase is +32. |

## Hint

`if c >= 'A' && c <= 'Z'`.

## Validate

```bash
make verify
```

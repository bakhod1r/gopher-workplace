# Word Wrap Width

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A terminal UI wraps text to a column width, but an off-by-one (`<` instead of
`<=`) breaks one character too early, so `"aa bb"` (exactly 5) is split at
width 5.

## Task

Fix the fit test between the markers in [wordwrap.go](wordwrap.go) so a line may
use the full width.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "the quick brown fox",10
Output: ["the quick" "brown fox"]
```

**Example 2:**

```
Input:  "aa bb cc",5
Output: ["aa bb" "cc"]
```

**Example 3:**

```
Input:  "a b",3
Output: ["a b"]
```

_Explanation:_ length 3 == width fits

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Fit condition** | New length `<= width` fits. |
| 2 | **Account for space** | `+1` for the joining space. |
| 3 | **Off-by-one** | `<` rejects an exactly-full line. |

## Hint

`if len(line)+1+len(w) <= width`.

## Validate

```bash
make verify
```

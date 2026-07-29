# Left Pad

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A report aligns numbers by zero-padding on the **left** (`42` → `00042`). The
code appends the pad on the right instead, so columns don't line up.

## Task

Fix the return between the markers in [padleft.go](padleft.go) to pad on the
left.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "42",5,'0'
Output: "00042"
```

**Example 2:**

```
Input:  "x",4,'.'
Output: "...x"
```

**Example 3:**

```
Input:  "toolong",3,'0'
Output: "toolong"
```

_Explanation:_ already >= width

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Pad side** | Left pad prepends fill. |
| 2 | **Width in runes** | Count runes, not bytes. |
| 3 | **No truncation** | Longer-than-width returns unchanged. |

## Hint

`return pad + s`.

## Validate

```bash
make verify
```

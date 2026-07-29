# Continue Skips Wrong Branch

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

`continue` skips the rest of the current iteration. The bug continues on the
positive elements — exactly the ones that should be summed — so only non-positive
values are added. The condition is inverted.

## Task

Fix [skipbug.go](skipbug.go) so the POSITIVE elements are summed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumPositive([1 -2 3 -4 5])
Output: 9
```

**Example 2:**

```
Input:  SumPositive([-1 -2])
Output: 0
```

**Example 3:**

```
Input:  SumPositive([1 2 3])
Output: 6
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **continue semantics** | It jumps to the next iteration, skipping code below. |
| 2 | **Guard polarity** | Skip the elements you DON'T want, not the ones you do. |
| 3 | **Filter-by-continue** | `if !want { continue }` then process. |

## Hint

Invert the guard: `if v <= 0 { continue }` so positives fall through to `sum += v`.

## Validate

```bash
make verify
```

# Missing Default Case

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

Every branch must produce a defined result. Falling out of the switch with an
empty return string leaves 3xx/1xx codes unlabeled; a `default` (or a non-empty
fallthrough return) fixes it.

## Task

Fix [classifybug.go](classifybug.go) so unknown classes return "unknown".

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Class(200)
Output: success
```

**Example 2:**

```
Input:  Class(301)
Output: unknown
```

**Example 3:**

```
Input:  Class(100)
Output: unknown
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Exhaustive branches** | Handle the else/default explicitly. |
| 2 | **default case** | Catches everything not listed. |
| 3 | **Defined output** | Never let a code fall through to "". |

## Hint

Return "unknown" for the uncovered classes — either a `default: return "unknown"` or change the trailing `return ""` to `return "unknown"`.

## Validate

```bash
make verify
```

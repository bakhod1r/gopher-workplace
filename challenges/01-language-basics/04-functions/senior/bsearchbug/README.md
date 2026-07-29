# Binary Search Bound

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

Binary search must shrink the interval every step. When `xs[mid] < target` the
new lower bound is `mid+1`; setting `lo = mid` can leave the interval unchanged
and spin (or time out).

## Task

Fix the low-bound update in [bsearchbug.go](bsearchbug.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IndexOf([1 3 5 7 9], 7)
Output: 3
```

**Example 2:**

```
Input:  IndexOf([1 3 5 7 9], 4)
Output: -1
```

**Example 3:**

```
Input:  IndexOf([1 3 5 7 9], 1)
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Interval invariant** | The search space must strictly shrink. |
| 2 | **Exclude mid** | mid was tested, so move to mid+1. |
| 3 | **Termination** | `lo = mid` can stall the loop. |

## Hint

Move past the tested element: `lo = mid + 1`.

## Validate

```bash
make verify
```

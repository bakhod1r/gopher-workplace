# Unsigned Underflow

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`have - sold` on `uint` wraps around when `sold > have`, producing a gigantic
number instead of 0. The zero value and unsigned wraparound combine into a
classic inventory bug.

## Task

Fix the single line between the markers in [stock.go](stock.go) so overselling
returns 0, never a wrapped value.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Remaining(10, 3)
Output: 7
```

**Example 2:**

```
Input:  Remaining(5, 5)
Output: 0
```

**Example 3:**

```
Input:  Remaining(2, 9)
Output: 0
```

_Explanation:_ Not a wrapped 18446744073709551609.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Unsigned wraparound** | `uint` has no negatives; `2-9` wraps to a huge value. |
| 2 | **Guard before subtract** | Compare `sold >= have` first. |
| 3 | **Zero value** | The clamped result is the `uint` zero. |

## Hint

Guard it: `if sold >= have { return 0 }` then `return have - sold`.

## Validate

```bash
make verify
```

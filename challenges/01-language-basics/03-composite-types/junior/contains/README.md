# Slice Contains

**Level:** junior
**Topic:** 01-language-basics → 03-composite-types

## Context

Checking membership in a small slice is a linear scan — return as soon as you
find it.

## Task

Implement `Contains(xs, target)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  xs=["a","b","c"], target="a"
Output: true
```

**Example 2:**

```
Input:  xs=["a","b","c"], target="z"
Output: false
```

**Example 3:**

```
Input:  xs=nil, target="x"
Output: false
```

_Explanation:_ Ranging a nil slice never executes the loop body.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Linear scan** | Compare each element. |
| 2 | **Early return** | Return true on the first match. |
| 3 | **Default false** | No match after the loop → false. |

## Hint

`for _, x := range xs { if x == target { return true } }; return false`.

## Validate

```bash
make verify
```

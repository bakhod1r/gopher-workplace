# Labeled Break Search

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

A labeled `break` exits an outer loop from inside a nested one — cleaner than a
found-flag checked at every level.

## Task

Implement `FindPairSum` in [firstpair.go](firstpair.go) using labeled break.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FindPairSum([1 2 3 4], 7)
Output: 2, 3, true (xs[2]+xs[3])
```

**Example 2:**

```
Input:  FindPairSum([1 2], 9)
Output: 0, 0, false
```

**Example 3:**

```
Input:  FindPairSum([3 4], 7)
Output: 0, 1, true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Labeled break** | `Outer: for ... { for ... { break Outer } }`. |
| 2 | **Nested search** | Scan all i<j pairs. |
| 3 | **Result flags** | Set i,j,ok then break out. |

## Hint

Label the outer loop; on a match set `i,j,ok` and `break Outer`.

## Validate

```bash
make verify
```

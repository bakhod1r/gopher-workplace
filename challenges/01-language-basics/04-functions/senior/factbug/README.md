# Recursion Base Case

**Level:** senior
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

Factorial's base case is `0! = 1`. Returning 0 makes the whole product
collapse to 0 because every level multiplies by the base.

## Task

Fix the base case in [factbug.go](factbug.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Fact(0)
Output: 1
```

**Example 2:**

```
Input:  Fact(5)
Output: 120
```

**Example 3:**

```
Input:  Fact(6)
Output: 720
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Recursion base case** | The terminating value must be correct. |
| 2 | **Multiplicative identity** | 1, not 0, is the neutral factor. |
| 3 | **Recursive product** | A wrong base propagates through every call. |

## Hint

Return `1` at `n == 0`.

## Validate

```bash
make verify
```

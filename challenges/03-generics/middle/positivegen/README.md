# All Positive

**Level:** middle  
**Topic:** 03-generics

## Context

A billing check rejects any batch containing a non-positive amount, and the operator needs to know which line failed.

## Task

Implement the stub(s) in [positivegen.go](positivegen.go):

1. Implement `AllPositive` and `FirstNonPositive`.
2. An empty slice is positive; `FirstNonPositive` returns `-1` when everything is positive.
3. Zero is not positive.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AllPositive([]int{1, 2})
Output: true
```

**Example 2:**

```
Input:  AllPositive([]int{1, 0})
Output: false
```

**Example 3:**

```
Input:  FirstNonPositive([]int{1, -1})
Output: 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Vacuous truth** | Reused from earlier: an empty slice has no counterexample. |
| 2 | **Boundary definitions** | "Positive" excludes zero — encode it with `<= 0`, not `< 0`. |
| 3 | **Narrow constraints on purpose** | Only signed sets make a non-positive value possible in the first place. |

## Hint

Test `v <= 0`, not `v < 0` — zero is not positive.

## Validate

```bash
make verify
```

# Spend A Budget

**Level:** senior
**Topic:** 04-error-handling

## Context

A request may perform a limited number of expensive operations. Exceeding the allowance is a failure, not a slowdown.

## Task

Implement `Spend` in [budgetsteps.go](budgetsteps.go):

1. Run steps in order while the budget allows, one unit each.
2. Return `ErrBudgetExceeded` without running a step once the budget is spent.
3. Return the first step failure immediately, leaving the rest unrun.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Spend(2, s1, s2)
Output: nil, both run
```

**Example 2:**

```
Input:  Spend(1, s1, s2)
Output: ErrBudgetExceeded after one
```

**Example 3:**

```
Input:  Spend(2, failing, s2)
Output: the failure, one run
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Resource budgets** | Bounded work per request. |
| 2 | **Failure precedence** | A step error outranks the budget check. |
| 3 | **Early exit accounting** | Unrun steps cost nothing. |

## Hint

Check the budget before each step, and stop the whole run on the first failure.

## Validate

```bash
make verify
```

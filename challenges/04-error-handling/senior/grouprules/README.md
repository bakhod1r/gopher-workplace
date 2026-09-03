# Validate With Rules

**Level:** senior
**Topic:** 04-error-handling

## Context

A submission is checked against a list of named rules. The response must name every rule that failed, in the order they were declared.

## Task

Implement `Check` in [grouprules.go](grouprules.go):

1. Run every rule, in order.
2. Wrap each failure as `"<name>: <err>"`.
3. Return the failures joined, or nil when all pass.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Check(v, passingRules)
Output: nil
```

**Example 2:**

```
Input:  Check(v, oneFailing)
Output: "len: too short"
```

**Example 3:**

```
Input:  Check(v, twoFailing)
Output: both, newline separated
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Named rules** | The rule name is the useful context. |
| 2 | **Run everything** | Independent rules all report. |
| 3 | **Deterministic order** | A slice preserves declaration order. |

## Hint

The rules arrive as a slice, so order is already fixed — the only decision is whether to stop early.

## Validate

```bash
make verify
```

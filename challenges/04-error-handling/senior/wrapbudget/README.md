# Cap The Chain

**Level:** senior
**Topic:** 04-error-handling

## Context

A deeply layered call stack wraps the same error at every level. The log line becomes unreadable, so annotation stops after a fixed depth.

## Task

Implement `Wrap` in [wrapbudget.go](wrapbudget.go):

1. Wrap `err` with `msg` when the chain is shorter than `max`.
2. Return `err` unchanged when the chain already has `max` or more links.
3. Return nil for a nil error.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Wrap(ErrBase, "a", 2)
Output: "a: base"
```

**Example 2:**

```
Input:  Wrap(twoDeep, "c", 2)
Output: twoDeep unchanged
```

**Example 3:**

```
Input:  Wrap(nil, "a", 2)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Chain depth** | Counting links before adding one. |
| 2 | **Bounded annotation** | Context has diminishing returns. |
| 3 | **Identity on refusal** | Returning the same value, not a copy. |

## Hint

Count the existing links first; only then decide whether another one is worth adding.

## Validate

```bash
make verify
```

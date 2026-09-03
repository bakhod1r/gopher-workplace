# Context Up To The Cause

**Level:** staff
**Topic:** 04-error-handling

## Context

A log line should show the annotations added above a known sentinel, and nothing below it.

## Task

Implement `Above` in [chainto.go](chainto.go):

1. Return the messages of each error in the chain until `target` is reached, exclusive.
2. Return nil when `target` is not in the chain.
3. Return an empty non-nil slice when `err` is `target` itself.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Above(fmt.Errorf("x: %w", ErrA), ErrA)
Output: ["x: a"]
```

**Example 2:**

```
Input:  Above(ErrA, ErrA)
Output: []
```

**Example 3:**

```
Input:  Above(ErrB, ErrA)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Partial traversal** | Stop at a known landmark. |
| 2 | **Two empty results** | Nil means absent, empty means found immediately. |
| 3 | **Chain walking** | One link at a time. |

## Hint

Absent and found-at-the-top are different answers — the tests distinguish nil from an empty slice.

## Validate

```bash
make verify
```

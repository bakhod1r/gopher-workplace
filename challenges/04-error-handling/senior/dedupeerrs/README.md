# Unique Failures

**Level:** senior
**Topic:** 04-error-handling

## Context

A fan-out over a hundred shards usually returns the same failure a hundred times. The summary reports each distinct message once, in first-seen order.

## Task

Implement `Unique` in [dedupeerrs.go](dedupeerrs.go):

1. Return one error per distinct message, in first-seen order.
2. Skip nil entries.
3. Return nil when nothing failed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Unique([]error{ErrA, ErrA, ErrB})
Output: [ErrA ErrB]
```

**Example 2:**

```
Input:  Unique([]error{nil})
Output: nil
```

**Example 3:**

```
Input:  Unique(nil)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deduplicating by message** | Distinct values may share a message. |
| 2 | **Insertion order** | A map alone loses order; pair it with a slice. |
| 3 | **Keeping the first occurrence** | The earliest error is the representative. |

## Hint

A `map[string]bool` decides membership; the output slice preserves the order.

## Validate

```bash
make verify
```

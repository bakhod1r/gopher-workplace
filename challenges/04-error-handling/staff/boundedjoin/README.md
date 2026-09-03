# Join With A Cap

**Level:** staff
**Topic:** 04-error-handling

## Context

A fan-out over ten thousand shards must report failures without producing an error whose message is a megabyte long.

## Task

Implement `Cap` in [boundedjoin.go](boundedjoin.go):

1. Join at most `max` of the non-nil errors, in order.
2. Append a final error reading `"and <n> more"` when entries were dropped.
3. Return nil when nothing failed.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Cap(2, ErrA, ErrB, ErrC)
Output: ErrA, ErrB and "and 1 more"
```

**Example 2:**

```
Input:  Cap(5, ErrA)
Output: just ErrA
```

**Example 3:**

```
Input:  Cap(2)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded output** | Error size is a resource too. |
| 2 | **Reporting the remainder** | A count preserves the scale. |
| 3 | **errors.Join with a summary** | The extra entry is just another error. |

## Hint

Count every failure, but only keep the first `max` — the summary line reports the difference.

## Validate

```bash
make verify
```

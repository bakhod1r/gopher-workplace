# Bounded Error Collector

**Level:** senior
**Topic:** 04-error-handling

## Context

An import job can produce millions of failures. The report keeps the first few and a count of the rest so the log stays usable.

## Task

Implement `Collector` in [errcollector.go](errcollector.go):

1. Store at most `limit` errors, in arrival order.
2. Count every non-nil error added, including dropped ones.
3. Return nil from `Err` when nothing was added, otherwise the stored errors joined.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  c.Add(ErrX); c.Count()
Output: 1
```

**Example 2:**

```
Input:  limit 2, add 5 errors
Output: 2 stored, count 5
```

**Example 3:**

```
Input:  empty collector Err()
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded accumulation** | Unbounded error lists are their own outage. |
| 2 | **Ignoring nil adds** | Success is not an entry. |
| 3 | **errors.Join over a bounded slice** | The stored subset is still matchable. |

## Hint

Three stubs share one struct. `Add` must count before deciding whether to store.

## Validate

```bash
make verify
```

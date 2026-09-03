# Remember The Failure

**Level:** senior
**Topic:** 04-error-handling

## Context

An expensive lookup is repeated for the same key many times a second. Both successes and failures are worth caching so a broken key is not retried constantly.

## Task

Implement `entry` in [errcache.go](errcache.go):

1. Call `load` at most once per key.
2. Return the cached value and error on later calls with the same key.
3. Cache failures as well as successes.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  two Get calls, same key
Output: load called once
```

**Example 2:**

```
Input:  failing key, two calls
Output: same error, load called once
```

**Example 3:**

```
Input:  different keys
Output: load called twice
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Negative caching** | Failures are results too. |
| 2 | **Struct-valued cache entries** | Value and error stored together. |
| 3 | **Comma-ok on the cache** | Presence decides whether to load. |

## Hint

Storing only the value makes a cached failure indistinguishable from a cached zero — cache both fields together.

## Validate

```bash
make verify
```

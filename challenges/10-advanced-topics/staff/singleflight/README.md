# One Miss, One Fetch, Many Waiters

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A cache entry expires and two hundred requests miss it at once. Two hundred identical database queries run, each building the same multi-megabyte result, and the database falls over.

## Task

Implement [singleflight.go](singleflight.go):

1. Run `fn` for `key` and return its result.
2. Concurrent callers for the same key must share one execution and all receive its result.
3. Different keys must not block each other.
4. After a call completes, the next `Do` for that key runs `fn` again — this is deduplication, not caching.

Replace the stub body in [singleflight.go](singleflight.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  32 goroutines calling g.Do("hot", fn)
Output: fn runs once, all get its result
```

**Example 2:**

```
Input:  two sequential Do calls
Output: fn runs twice
```

_Explanation:_ No caching across completed calls.

**Example 3:**

```
Input:  8 goroutines on 8 keys
Output: fn runs 8 times
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Deduplication vs caching** | One shares an in-flight call; the other stores a finished result. |
| 2 | **WaitGroup as a one-shot broadcast** | Waiters block until the owner calls Done. |
| 3 | **Lock scope** | The map is guarded; `fn` runs outside the lock or every key serialises. |
| 4 | **Cleanup after completion** | Removing the entry is what makes the next call run again. |

## Hint

The first caller for a key owns the call. Everyone else waits on it and reads its result.

## Validate

```bash
make verify
```

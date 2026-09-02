# Single Flight

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A cache stampede hit the origin: a thousand concurrent misses on the same key produced a thousand identical fetches.

## Task

Implement the stub(s) in [singleflight.go](singleflight.go):

1. Implement `Do` on `*Group`, so concurrent calls for the same key share one execution of the loader.
2. Every caller must receive the shared result; distinct keys must not block each other.
3. Constraint: race-free under `-race`, and the loader must run exactly once per key per in-flight window.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  100 concurrent Do calls for one key
Output: the loader runs once, all 100 get the value
```

**Example 2:**

```
Input:  two different keys concurrently
Output: two loader runs, no serialisation between them
```

**Example 3:**

```
Input:  a second Do after the first completes
Output: the loader runs again
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Duplicate suppression** | One in-flight call per key; the rest wait on it. |
| 2 | **sync.WaitGroup as a broadcast** | `wg.Wait()` releases every waiter when the leader finishes. |
| 3 | **Lock scope discipline** | The map lock must not be held while the loader runs. |

## Hint

Store a `*call` in the map, release the map lock, then `wg.Wait()` on it.

## Validate

```bash
make verify
```

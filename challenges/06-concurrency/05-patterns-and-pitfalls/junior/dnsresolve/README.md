# Concurrent DNS Resolve

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

Service discovery resolves a few hundred hostnames on every config reload.
Resolving them one after another takes seconds; resolving them concurrently
takes as long as the slowest lookup — provided the shared result map is
written safely, because concurrent map writes crash the process outright.

## Task

Implement `ResolveAll` in [dnsresolve.go](dnsresolve.go) so that:

1. It creates the result map up front and starts one goroutine per host.
2. Each goroutine calls `lookup(host)` *without* holding the lock.
3. It then takes the mutex to write the entry, and the map is returned after `wg.Wait()`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ResolveAll([]string{"1"}, lookup)
Output: map[1:10.0.0.1]
```

**Example 2:**

```
Input:  ResolveAll([]string{"1", "2"}, lookup)
Output: map[1:10.0.0.1 2:10.0.0.2]
```

**Example 3:**

```
Input:  ResolveAll(nil, lookup)
Output: empty map
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Concurrent map writes** | Unsynchronised map writes are a fatal runtime error, not a recoverable race. |
| 2 | **Small critical sections** | Do the slow lookup outside the lock, hold it only for the assignment. |
| 3 | **Map identity** | Every goroutine writes into the same map value — maps are reference types. |

## Hint

Resolve first, lock second: `addr := lookup(host)` then `mu.Lock()`,
`addrs[host] = addr`, `mu.Unlock()`.

## Validate

```bash
make verify
```

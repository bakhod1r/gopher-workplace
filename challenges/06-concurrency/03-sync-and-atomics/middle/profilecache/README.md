# Read-Mostly Profile Cache

**Level:** middle  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

Every authenticated request needs the caller's profile. Fetching it from the identity service on each request is far too slow, so profiles are cached in process. Reads dominate; writes happen only when a profile is refreshed or invalidated. The cache also has to report its hit rate without turning every read into an exclusive lock.

## Task

Implement the exported function(s) in [profilecache.go](profilecache.go) so that:

1. `Put` takes the write lock and stores the profile.
2. `Get` takes the read lock, looks the profile up, and records a hit or a miss with the atomic counters.
3. `Invalidate` takes the write lock, deletes the entry, and reports whether one existed.
4. `Stats` returns the two counters.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  c.Put("u1", "ada"); c.Get("u1")
Output: "ada", true
```

**Example 2:**

```
Input:  NewCache().Get("nobody")
Output: "", false
```

**Example 3:**

```
Input:  c.Put("u1", "ada"); c.Get("u1"); c.Get("u2"); c.Stats()
Output: 1, 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.RWMutex** | Many `RLock` holders at once; `Lock` excludes all of them. |
| 2 | **Atomic counters beside a lock** | Stats update under the read lock because `atomic.Int64` is already safe on its own. |
| 3 | **Release before you count** | Do the atomic add outside the critical section — it does not need the map. |

## Hint

`Get` should hold `RLock` only around the map lookup. Increment `hits`/`misses` with `Add(1)` after unlocking.

## Validate

```bash
make verify
go test -race ./...
```

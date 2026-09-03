# Catalogue Fetch Deduplication

**Level:** middle
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A product page renders from a slow catalogue service. During a traffic spike, hundreds of goroutines ask for the same product at once — the cache must collapse them into a single upstream call per key and hand everyone the same answer.

## Task

Implement the stubbed functions in [dedupefetch.go](dedupefetch.go) so that:

1. `NewFetcher` returns a Fetcher with its map initialised.
2. `Fetch` returns the cached value, calling `load` **exactly once per key**.
3. Concurrent callers for the same key must collapse into a single `load` call.
4. The map lookup must be protected, but `load` must not run under the map lock.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  f.Fetch("sku-1", load)
Output: "product:sku-1"
```

**Example 2:**

```
Input:  f.Fetch("sku-1", load) again
Output: "product:sku-1", load not called
```

**Example 3:**

```
Input:  f.Fetch("sku-2", load)
Output: "product:sku-2"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | `sync.Once` per key | One `Once` per entry turns "exactly once" from a race into a guarantee, whatever the scheduling. |
| 2 | Narrow critical section | Hold the map lock only to find or create the entry; run the slow `load` outside it. |
| 3 | Check-then-act on a map | `if _, ok := m[k]; !ok { m[k] = ... }` is only safe while the lock is held for the whole sequence. |
| 4 | Happens-before from `Once` | Every `Do` caller observes the writes made inside the function — reading `e.value` afterwards needs no extra lock. |

## Hint

Two phases: (1) under the mutex, get-or-create the `*entry`; (2) outside the mutex, `e.once.Do(func(){ e.value = load(key) })` and return `e.value`.

## Validate

```bash
make verify
```

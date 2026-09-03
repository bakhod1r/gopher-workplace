# Copy-on-Write Rollout Plan

**Level:** middle  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

Feature-flag percentages are consulted on every single request and changed a handful of times a day. Even an `RWMutex` costs an atomic read-modify-write per lookup. Because the data is tiny and immutable once published, writers can copy the whole map, mutate the copy, and swap the pointer — readers then do a single atomic load and no locking at all.

## Task

Implement the exported function(s) in [rolloutplan.go](rolloutplan.go) so that:

1. `NewPlan` stores a *copy* of `initial` with every percentage clamped into `0..100`, so a later mutation of the caller's map cannot change the plan.
2. `Percent` loads the current snapshot and returns the feature's value, or `0` when unknown, without taking any lock.
3. `SetPercent` holds the writer mutex, copies the current map, sets the clamped value, and publishes the copy.
4. `Features` returns the names in the current snapshot, sorted.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  NewPlan(map[string]int{"checkout": 25}).Percent("checkout")
Output: 25
```

**Example 2:**

```
Input:  p := NewPlan(nil); p.SetPercent("checkout", 300); p.Percent("checkout")
Output: 100
```

**Example 3:**

```
Input:  p.SetPercent("b", 1); p.SetPercent("a", 1); p.Features()
Output: ["a" "b"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Copy-on-write** | A published snapshot is immutable; a change means a new map, never an edit in place. |
| 2 | **atomic.Pointer[T]** | `Load`/`Store` publish a whole data structure with one atomic operation. |
| 3 | **Writer-only mutex** | The lock exists to make read-copy-publish atomic *between writers*; readers never touch it. |

## Hint

`SetPercent` must read the old map, build a brand-new one, and `Store` it — all while holding `p.writers`. Never write into the map a reader might be holding.

## Validate

```bash
make verify
go test -race ./...
```

# Copy On Write

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A config snapshot is read on every request and updated rarely. Locking readers is wasteful; the config is now swapped atomically instead.

## Task

Implement the stub(s) in [copyonwrite.go](copyonwrite.go):

1. Implement `Load` and `Store` on `*Config` using `atomic.Value`.
2. Implement `Update`, which applies a `Mutator` to a *copy* of the current snapshot and stores the new one.
3. Constraint: readers must never see a partially updated snapshot, and `-race` must be clean with concurrent readers and writers.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Store a snapshot, then Load
Output: the same values
```

**Example 2:**

```
Input:  Update adding a key
Output: a new snapshot; the old one is unchanged
```

**Example 3:**

```
Input:  concurrent readers during updates
Output: every read sees a complete snapshot
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Copy-on-write** | Readers are lock-free; writers pay the copy. |
| 2 | **atomic.Value** | One atomic pointer swap publishes a whole immutable snapshot. |
| 3 | **Immutability discipline** | A stored snapshot must never be mutated in place. |

## Hint

Copy the map inside `Update`, mutate the copy, then `Store` it — never write into the loaded map.

## Validate

```bash
make verify
```

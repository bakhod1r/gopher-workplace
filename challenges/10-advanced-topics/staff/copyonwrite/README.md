# Publish A New Map, Never Edit The Old One

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A read-mostly routing table is guarded by an RWMutex. Reads dominate by a thousand to one, and the read lock's atomic traffic is the top entry in the profile.

## Task

Implement [copyonwrite.go](copyonwrite.go):

1. Publish a new snapshot with `key` set to `val`.
2. A published map must never be modified — readers may still be holding it.
3. Serialise writers with the mutex; readers take no lock.
4. Correct with concurrent readers and writers.

Replace the stub body in [copyonwrite.go](copyonwrite.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s.Set("a", 1); s.Get("a")
Output: 1, true
```

**Example 2:**

```
Input:  a snapshot loaded before a Set
Output: unchanged afterwards
```

**Example 3:**

```
Input:  4 writers x 200 sets, 8 readers
Output: no torn or lost state
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Copy-on-write** | Writers pay a full copy so readers pay nothing. |
| 2 | **Atomic pointer publication** | One word swap replaces the whole map indivisibly. |
| 3 | **Immutable after publication** | The published map is shared with every reader, forever. |
| 4 | **The writer lock** | It serialises the read-copy-swap sequence, which is not atomic on its own. |

## Hint

Load, copy, modify the copy, store. The lock is around all four.

## Validate

```bash
make verify
```

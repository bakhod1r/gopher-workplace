# Dedup Stream

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

An event stream contains duplicates. Deduplication must run over 100M events without buffering them, using a bounded seen-set.

## Task

Implement the stub(s) in [dedupstream.go](dedupstream.go):

1. Implement `Seen` on `*ExactSet` — report whether the id was seen before and record it.
2. Implement `Seen` on `*WindowSet` — remember only the last `N` ids (a sliding window).
3. Implement `Dedup`, which streams a `Source` through a `SeenSet` and returns the number of unique events.
4. Constraint: `WindowSet` memory must stay bounded by `N` regardless of stream length.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Dedup over [a a b] with an ExactSet
Output: 2
```

**Example 2:**

```
Input:  WindowSet{N: 1} over [a b a]
Output: 3 (a fell out of the window)
```

**Example 3:**

```
Input:  an empty stream
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded deduplication** | Exact dedup is O(unique) memory; a window trades accuracy for a ceiling. |
| 2 | **Set + queue** | Reused: a map for membership, a slice for eviction order. |
| 3 | **Streaming** | Reused: never materialise the input. |

## Hint

`WindowSet` evicts the oldest id from both the map and the queue when it exceeds `N`.

## Validate

```bash
make verify
```

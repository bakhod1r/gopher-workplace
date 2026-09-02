# Merge Many Streams

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A query fans out to dozens of sorted shards. Their results must merge into one sorted stream without buffering everything.

## Task

Implement the stub(s) in [mergemany.go](mergemany.go):

1. Implement `Next` on `*SortedFeed`.
2. Implement `MergeAll`, which merges any number of ascending feeds into one ascending slice.
3. Constraint: memory must be O(number of feeds), not O(total elements) — pick the minimum head each step, do not concatenate and sort.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  merge [1 4], [2 5], [3]
Output: [1 2 3 4 5]
```

**Example 2:**

```
Input:  merge with an empty feed
Output: the others still merge
```

**Example 3:**

```
Input:  merge nothing
Output: empty
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **K-way merge** | Only the heads of the feeds are held at once. |
| 2 | **Streaming discipline** | Reused: the source is pulled, never materialised. |
| 3 | **Stable tie-breaking** | Equal heads are taken from the earliest feed, keeping output deterministic. |

## Hint

Keep a head slice with one entry per feed; refill only the slot you consumed.

## Validate

```bash
make verify
```

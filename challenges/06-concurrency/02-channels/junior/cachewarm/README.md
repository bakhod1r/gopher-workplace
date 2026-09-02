# Next Cache Key

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The cache warmer pulls at most one key id from the eviction feed per tick.
It has to tell apart a key id of 0 from a feed that has already finished.

## Task

Implement `NextKey` in [cachewarm.go](cachewarm.go) so that:

1. It performs exactly one receive on `feed`.
2. It returns the key id and `true` when the feed was still open.
3. It returns `0, false` when the feed is closed and drained.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NextKey(5)
Output: 5, true
```

**Example 2:**

```
Input:  NextKey() // closed, empty
Output: 0, false
```

**Example 3:**

```
Input:  NextKey(0)
Output: 0, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Comma-ok receive** | `v, ok := <-ch` — `ok` is false only for a closed, empty channel. |
| 2 | **Zero value** | A closed channel yields the element type's zero value forever. |
| 3 | **Multiple returns** | Value plus validity is the Go way to signal "nothing here". |

## Hint

One line: `id, ok := <-feed`. The second result is exactly the boolean you
need.

## Validate

```bash
make verify
```

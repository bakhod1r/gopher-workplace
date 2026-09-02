# sort.Interface

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A leaderboard sorts players by score, highest first, with names breaking ties.

## Task

Implement the stub(s) in [sortslice.go](sortslice.go):

1. Implement `Len`, `Less`, and `Swap` on `ByScore` so `sort.Sort` orders players by descending score, then ascending name.
2. Implement `TopN`, which returns the first n names after sorting (fewer if the list is shorter).
3. Sort in place — do not copy the slice inside `Less` or `Swap`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  sorted names for scores {a:1, b:3}
Output: ["b", "a"]
```

**Example 2:**

```
Input:  tie between "ann" (5) and "bob" (5)
Output: ["ann", "bob"]
```

**Example 3:**

```
Input:  TopN(players, 10) with 2 players
Output: both names
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sort.Interface** | `sort.Sort` drives any type with `Len/Less/Swap`. |
| 2 | **Multi-key ordering** | Compare the secondary key only when the primary ties. |
| 3 | **Stable results** | Reused: deterministic output makes the test meaningful. |

## Hint

`Less(i, j)` means "i sorts before j" — for descending scores that is `>`.

## Validate

```bash
make verify
```

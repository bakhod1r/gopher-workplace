# Index Of Max

**Level:** junior  
**Topic:** 03-generics

## Context

A leaderboard highlights the winning row. Knowing the value is not enough — the row's position drives the rendering.

## Task

Implement the stub(s) in [indexmax.go](indexmax.go):

1. Implement `IndexOfMax`, returning the index of the largest element.
2. On a tie return the earliest index; return `-1` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  IndexOfMax([]int{1, 9, 9})
Output: 1
```

**Example 2:**

```
Input:  IndexOfMax([]string{"c", "a"})
Output: 0
```

**Example 3:**

```
Input:  IndexOfMax([]int{})
Output: -1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Tracking an index, not a value** | Comparing `s[i] > s[best]` keeps the index authoritative. |
| 3 | **Sentinel `-1`** | Reused from earlier: the conventional "no index" answer. |

## Hint

Keep `best` as an index and compare `s[i]` against `s[best]`.

## Validate

```bash
make verify
```

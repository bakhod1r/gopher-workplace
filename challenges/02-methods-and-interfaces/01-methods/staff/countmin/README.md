# Count-Min Sketch

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A count-min sketch estimates frequencies in fixed space. Each item increments
one counter per row; the estimate is the *minimum* across rows, because
collisions can only inflate a counter, never deflate it.

## Task

Implement `Add` and `Count` on `*Sketch` in [countmin.go](countmin.go):

1. `Add` increments `s.row1[h1(item)]` and `s.row2[h2(item)]`.
2. `Count` returns the smaller of those two counters.
3. The empty-string guards are already written — leave them.

**Constraint (staff):** the estimate must never fall below the true count across 5,000 additions, and counting must not allocate.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  Add("apple") twice; Count("apple")
Output: 2
```

**Example 2:**

```
Input:  Add("bat"); Count("bat")
Output: 1
```

**Example 3:**

```
Input:  Count("never added") where both its buckets are untouched
Output: 0
```

_Explanation:_ an unseen item reads whatever its buckets hold — 0 when nothing collided.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Minimum, not sum or maximum** | Every row over-counts on collision, so the min is the tightest upper bound. |
| 2 | **Never underestimates** | The true count is always ≤ the estimate. |
| 3 | **`min` builtin** | Go 1.21+ has `min(a, b)` for ordered types — no helper needed. |

## Hint

`Count` is `return min(s.row1[h1(item)], s.row2[h2(item)])`. Summing the rows
double-counts; taking the max amplifies collisions instead of damping them.

## Validate

```bash
make verify
```

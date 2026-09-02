# Top-K Stream

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A dashboard needs the top 10 scores from a 100M-row stream. Sorting the whole stream is not an option.

## Task

Implement the stub(s) in [topkstream.go](topkstream.go):

1. Implement `Add` and `Result` on `*TopK`, keeping only the k largest values seen.
2. Implement `Stream`, which folds a `Source` through a `TopKAgg` and returns the results in descending order.
3. Constraint: memory must be O(k), not O(n) — the test asserts the internal slice never exceeds k.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  k 2 over [1 5 3]
Output: [5 3]
```

**Example 2:**

```
Input:  k 3 over [1 2]
Output: [2 1]
```

**Example 3:**

```
Input:  k 0
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded top-k** | Keep k items, not n. |
| 2 | **Insertion into a sorted window** | O(k) per element is fine when k is small and fixed. |
| 3 | **Streaming aggregation** | Reused: fold, never buffer. |

## Hint

Keep the window sorted descending; a new value only matters when it beats the smallest kept one.

## Validate

```bash
make verify
```

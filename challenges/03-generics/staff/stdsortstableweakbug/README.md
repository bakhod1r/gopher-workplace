# A Comparator That Cannot Decide

**Level:** staff  
**Topic:** 03-generics

## Context

A scheduler sorts a work queue by priority and is supposed to preserve arrival order inside each priority. Same-priority work keeps coming out backwards, and on big queues the order is not even sorted.

## Task

Fix the single planted bug in [stdsortstableweakbug.go](stdsortstableweakbug.go):

1. Find and fix the single bug so the comparator reports a real tie as a tie.
2. The sort must stay ascending by `Pri` and stable within a priority.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  SortByPriority([{a 2} {b 1} {c 2}])
Output: [{b 1} {a 2} {c 2}]
```

**Example 2:**

```
Input:  three tasks all at Pri 1
Output: input order preserved
```

**Example 3:**

```
Input:  input slice after the call
Output: unchanged
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Strict weak ordering** | A comparator must be antisymmetric: if `f(a,b) < 0` then `f(b,a) > 0`, and equals must return 0. |
| 2 | **Stability needs equality** | A stable sort keeps equal elements in order — it cannot, if nothing is ever equal. |

## Hint

Feed the comparator the same task twice. What does it say?

## Validate

```bash
make verify
```

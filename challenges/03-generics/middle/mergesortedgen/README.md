# Merge Sorted

**Level:** middle  
**Topic:** 03-generics

## Context

Two already-sorted result pages must be shown as one ordered list without re-sorting everything.

## Task

Implement the stub(s) in [mergesortedgen.go](mergesortedgen.go):

1. Implement `Merge`, returning all elements of both sorted inputs in order.
2. On ties take the element from `a` first.
3. Do not sort — a single linear pass is required.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Merge([]int{1,3}, []int{2,4})
Output: []int{1,2,3,4}
```

**Example 2:**

```
Input:  Merge([]int{}, []int{1})
Output: []int{1}
```

**Example 3:**

```
Input:  Merge([]int{1}, []int{1})
Output: []int{1,1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Two-pointer merge** | The core of merge sort: advance whichever side has the smaller head. |
| 2 | **Tie handling defines stability** | Taking from `a` on ties is what keeps the merge stable. |
| 3 | **`cmp.Ordered`** | The stdlib constraint for `<`, `<=`, `>`, `>=`. |

## Hint

Compare `b[j] < a[i]` — the strict comparison is what makes ties favour `a`.

## Validate

```bash
make verify
```

# Binary Search

**Level:** junior  
**Topic:** 03-generics

## Context

A sorted index of IDs is searched thousands of times per request. A linear scan is no longer good enough.

## Task

Implement the stub(s) in [slicesbinsearch.go](slicesbinsearch.go):

1. Implement `Find` using `slices.BinarySearch`.
2. The input is assumed sorted; the second return value reports whether `v` was found.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Find([]int{1, 3, 5}, 3)
Output: 1, true
```

**Example 2:**

```
Input:  Find([]int{1, 3, 5}, 4)
Output: 2, false
```

**Example 3:**

```
Input:  Find([]int{}, 1)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.BinarySearch`** | Returns `(index, found)`; the index is where `v` is, or where it would be inserted. |
| 2 | **Sorted precondition** | Binary search on unsorted input returns nonsense, not an error. |
| 3 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<` and `>`. |

## Hint

The insertion point comes back even when the value is missing — that is the useful part.

## Validate

```bash
make verify
```

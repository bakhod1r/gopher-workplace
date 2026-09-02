# Binary Search With Comparison

**Level:** middle  
**Topic:** 03-generics

## Context

A large table sorted by ID is probed constantly, and the rows are structs with no natural ordering.

## Task

Implement the stub(s) in [slicesbsearchfunc.go](slicesbsearchfunc.go):

1. Implement `FindByID` using `slices.BinarySearchFunc`.
2. The comparison receives an element and the target — note the asymmetric signature.
3. The rows are assumed sorted by ID.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FindByID(rows, 3)
Output: index, true
```

**Example 2:**

```
Input:  FindByID(rows, 4)
Output: insertion point, false
```

**Example 3:**

```
Input:  FindByID(nil, 1)
Output: 0, false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.BinarySearchFunc`** | The comparison takes `(element, target)` — the target type may differ from the element type. |
| 2 | **Comparing against a key** | You never need to build a fake `Row` just to search. |
| 3 | **Sorted precondition** | Unsorted input gives a wrong answer, not an error. |

## Hint

The comparison's second parameter is the target, not another element.

## Validate

```bash
make verify
```

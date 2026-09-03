# Do These Two Slices Share Storage

**Level:** junior
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A helper documents that it may return its input or a fresh slice. A caller needs to know which happened before it mutates the result.

## Task

Implement [samearray.go](samearray.go):

1. Report whether `a` and `b` begin at the same element of the same array.
2. Any empty operand reports false.

Replace the stub body in [samearray.go](samearray.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  s := []int{1,2,3}; SameArray(s, s[:2])
Output: true
```

_Explanation:_ Same start, different length.

**Example 2:**

```
Input:  SameArray(s, s[1:])
Output: false
```

_Explanation:_ The start differs.

**Example 3:**

```
Input:  SameArray(a, b) after copy
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **unsafe.SliceData** | The slice's data pointer, comparable with == |
| 2 | **Slices are not comparable** | `a == b` does not compile for slices; only `== nil` is allowed. |
| 3 | **Aliasing detection** | Shared storage is what makes one write visible through another slice. |

## Hint

Two pointers, one comparison — plus a guard for empties.

## Validate

```bash
make verify
```

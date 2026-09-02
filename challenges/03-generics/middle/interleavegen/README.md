# Interleave

**Level:** middle  
**Topic:** 03-generics

## Context

A comparison view shows old and new lines alternately, then whatever is left of the longer side.

## Task

Implement the stub(s) in [interleavegen.go](interleavegen.go):

1. Implement `Interleave`, alternating elements starting with `a`.
2. Append the remainder of the longer slice once the shorter runs out.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Interleave([]int{1,3}, []int{2,4})
Output: []int{1,2,3,4}
```

**Example 2:**

```
Input:  Interleave([]int{1,2,3}, []int{9})
Output: []int{1,9,2,3}
```

**Example 3:**

```
Input:  Interleave([]int{}, []int{1})
Output: []int{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Paired traversal** | One index walks both slices while both have elements. |
| 2 | **Draining the remainder** | Both remainder appends are safe; one is always empty. |
| 3 | **Multi-value append** | `append(out, a[i], b[i])` adds two elements in one call. |

## Hint

One shared index, then two remainder appends.

## Validate

```bash
make verify
```

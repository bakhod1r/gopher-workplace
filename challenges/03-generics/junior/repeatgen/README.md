# Repeat Slice

**Level:** junior  
**Topic:** 03-generics

## Context

A test fixture builder needs a payload repeated a few times to exercise pagination.

## Task

Implement the stub(s) in [repeatgen.go](repeatgen.go):

1. Implement `Repeat`, returning `s` concatenated with itself `n` times.
2. Return an empty (non-nil) slice when `n` is zero or negative.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Repeat([]int{1, 2}, 2)
Output: []int{1, 2, 1, 2}
```

**Example 2:**

```
Input:  Repeat([]string{"a"}, 3)
Output: []string{"a", "a", "a"}
```

**Example 3:**

```
Input:  Repeat([]int{1}, 0)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Variadic append** | `append(out, s...)` splices a whole slice in one call, for any element type. |
| 2 | **Capacity hints** | `len(s)*n` is the exact final length — pre-allocating avoids repeated regrowth. |
| 3 | **Slices of a type parameter** | `[]T` behaves like any slice: `len`, `range`, `append` all work. |

## Hint

`append(out, s...)` inside the repeat loop does the whole copy.

## Validate

```bash
make verify
```

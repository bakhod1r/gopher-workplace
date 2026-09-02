# Sliding Window

**Level:** middle  
**Topic:** 03-generics

## Context

A smoothing filter averages every three consecutive samples, so it needs the windows before it can average them.

## Task

Implement the stub(s) in [windowgen.go](windowgen.go):

1. Implement `Windows`, returning each consecutive window of `n` elements.
2. Return an empty result when `n <= 0` or `n` exceeds the slice length.
3. Each window must be independent of the input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Windows([]int{1,2,3}, 2)
Output: [][]int{{1,2},{2,3}}
```

**Example 2:**

```
Input:  Windows([]int{1,2}, 3)
Output: [][]int{}
```

**Example 3:**

```
Input:  Windows([]int{1,2}, 2)
Output: [][]int{{1,2}}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Overlapping versus chunking** | Windows overlap by `n-1`; chunks do not overlap at all. |
| 2 | **No aliasing** | Return fresh slices; sub-slices of the input share its backing array. |
| 3 | **Loop bound** | `i+n <= len(s)` is the condition that keeps the last window full. |

## Hint

There are `len(s)-n+1` windows; the loop condition `i+n <= len(s)` gives you exactly that.

## Validate

```bash
make verify
```

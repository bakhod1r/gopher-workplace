# Arrays Copy, Slices Point

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`[8]int` and `[]int` look almost identical and behave nothing alike. An array is a value: passing it copies every element, and the callee cannot touch the original. A slice is a three-word view: passing it copies the view and shares the data.

## Task

Implement all three in [arrayvsslice.go](arrayvsslice.go):

1. `SumBlock` sums a `Block`, and `SumSlice` sums a `[]int`.
2. `ZeroBlock` returns an all-zero block while leaving its argument untouched.
3. None of them may allocate.

## Examples

**Example 1:**

```
Input:  SumBlock(Block{1, 2, 3})
Output: 6
```

**Example 2:**

```
Input:  b = Block{1,2,3}; ZeroBlock(b)
Output: zero Block returned, b still {1 2 3 ...}
```

**Example 3:**

```
Input:  s := b[:]; s[0] = 99
Output: b[0] is 99 — the array is shared
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Array size is part of its type** | `[8]int` and `[9]int` are different types, and both copy on assignment. |
| 2 | **Copy cost scales with the array** | Eight words is fine; a `[4096]byte` parameter is a memcpy per call. |
| 3 | **`b[:]` shares, it does not copy** | Slicing an array creates a view onto that exact array. |

## Topics used again

Arrays, slices, value semantics, `range`.

## Hint

`ZeroBlock` can return the zero value of the type rather than looping.

## Validate

```bash
make verify
```

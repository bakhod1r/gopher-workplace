# Grow Before Appending

**Level:** middle  
**Topic:** 03-generics

## Context

A hot loop appends thousands of records to an existing slice, and the repeated regrowth shows up in profiles.

## Task

Implement the stub(s) in [slicesgrowgen.go](slicesgrowgen.go):

1. Implement `Collect` using `slices.Grow` before appending.
2. The result must contain the original elements followed by the new ones.
3. After the call the capacity must be at least the final length.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Collect([]int{1}, 2, 3)
Output: []int{1,2,3}
```

**Example 2:**

```
Input:  cap after the call
Output: >= 3
```

**Example 3:**

```
Input:  Collect(nil, 1)
Output: []int{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Grow`** | Reserves room for `n` more elements, growing at most once. |
| 2 | **Amortised versus reserved** | Append doubles as it grows; `Grow` makes it a single allocation. |
| 3 | **Length is unchanged** | `Grow` changes capacity only — the elements still come from `append`. |

## Hint

`Grow` returns a slice with room; it does not change the length.

## Validate

```bash
make verify
```

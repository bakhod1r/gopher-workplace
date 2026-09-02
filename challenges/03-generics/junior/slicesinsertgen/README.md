# Insert

**Level:** junior  
**Topic:** 03-generics

## Context

A playlist inserts a track at the position the user dropped it. An out-of-range drop must be ignored, not fatal.

## Task

Implement the stub(s) in [slicesinsertgen.go](slicesinsertgen.go):

1. Implement `InsertAt` using `slices.Insert`.
2. Return `s` unchanged for an out-of-range index; inserting at `len(s)` is valid and appends.
3. Do not modify the caller's slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  InsertAt([]int{1, 3}, 1, 2)
Output: []int{1, 2, 3}
```

**Example 2:**

```
Input:  InsertAt([]int{1}, 1, 2)
Output: []int{1, 2}
```

**Example 3:**

```
Input:  InsertAt([]int{1}, 5, 2)
Output: []int{1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Insert`** | `Insert(s, i, v...)` returns the grown slice and panics on a bad index. |
| 2 | **Insert at `len(s)`** | That position is legal — it means append. |
| 3 | **Cloning first** | Reused from earlier: clone before an in-place helper when the caller must keep its data. |

## Hint

`slices.Insert` panics on a bad index, so validate before calling. Valid indexes run `0..len(s)`.

## Validate

```bash
make verify
```

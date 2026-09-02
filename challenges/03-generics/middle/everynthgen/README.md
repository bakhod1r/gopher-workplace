# Every Nth

**Level:** middle  
**Topic:** 03-generics

## Context

A sparkline downsamples a long series to a handful of points, deterministically, so two renders agree.

## Task

Implement the stub(s) in [everynthgen.go](everynthgen.go):

1. Implement `EveryNth`, returning the elements at indexes `0, n, 2n, ...`.
2. Return an empty (non-nil) result when `n <= 0`.
3. `n == 1` returns a copy of the whole slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  EveryNth([]int{1,2,3,4,5}, 2)
Output: []int{1,3,5}
```

**Example 2:**

```
Input:  EveryNth([]int{1,2}, 1)
Output: []int{1,2}
```

**Example 3:**

```
Input:  EveryNth([]int{1,2}, 0)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Strided loops** | `i += n` walks the slice with a stride instead of one at a time. |
| 2 | **Guarding the stride** | A stride of 0 would loop forever — the guard is not optional. |
| 3 | **Deterministic sampling** | Index-based downsampling is reproducible; random sampling is not. |

## Hint

Guard `n <= 0` before the loop, or a stride of zero hangs.

## Validate

```bash
make verify
```

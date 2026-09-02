# Repeat From Stdlib

**Level:** middle  
**Topic:** 03-generics

## Context

A test fixture needs a payload repeated a configurable number of times, and the count comes from a table that may hold zero.

## Task

Implement the stub(s) in [slicesrepeatstd.go](slicesrepeatstd.go):

1. Implement `Tile` using `slices.Repeat`.
2. Return an empty (non-nil) slice when `n <= 0` — `slices.Repeat` panics on a negative count.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Tile([]int{1,2}, 2)
Output: []int{1,2,1,2}
```

**Example 2:**

```
Input:  Tile([]int{1}, 0)
Output: []int{}
```

**Example 3:**

```
Input:  Tile([]int{1}, -1)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`slices.Repeat`** | Concatenates `count` copies; it panics for a negative count. |
| 2 | **Wrapping panicking helpers** | Reused judgement: add the guard your callers need. |
| 3 | **Overflow check** | `Repeat` also panics when the result length would overflow. |

## Hint

Guard the count yourself — the stdlib panics rather than returning empty.

## Validate

```bash
make verify
```

# Rotate

**Level:** middle  
**Topic:** 03-generics

## Context

A carousel scrolls by an offset the user can drag in either direction, past the ends and back again.

## Task

Implement the stub(s) in [rotategen.go](rotategen.go):

1. Implement `Rotate`, returning a new slice rotated left by `k`.
2. Negative `k` rotates right; `k` larger than the length wraps around.
3. Leave the input unmodified.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Rotate([]int{1,2,3}, 1)
Output: []int{2,3,1}
```

**Example 2:**

```
Input:  Rotate([]int{1,2,3}, -1)
Output: []int{3,1,2}
```

**Example 3:**

```
Input:  Rotate([]int{1,2,3}, 4)
Output: []int{2,3,1}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Normalising a modulus** | `((k % n) + n) % n` maps any integer into `[0, n)`, unlike `%` alone. |
| 2 | **Go's `%` keeps the sign** | `-1 % 3` is `-1` in Go, so a bare modulus is not enough. |
| 3 | **No aliasing** | Return fresh slices; sub-slices of the input share its backing array. |

## Hint

Normalise `k` into `[0, len(s))` first, then append the two halves.

## Validate

```bash
make verify
```

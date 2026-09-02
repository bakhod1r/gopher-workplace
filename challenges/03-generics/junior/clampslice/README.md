# Clamp All

**Level:** junior  
**Topic:** 03-generics

## Context

Sensor readings occasionally spike outside the physically possible range. The plot clamps them instead of dropping the sample.

## Task

Implement the stub(s) in [clampslice.go](clampslice.go):

1. Implement `ClampAll`, returning a new slice with each element limited to `[lo, hi]`.
2. Leave the input slice unmodified.
3. Return an empty (non-nil) slice for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ClampAll([]int{-1, 2, 9}, 0, 3)
Output: []int{0, 2, 3}
```

**Example 2:**

```
Input:  ClampAll([]float64{5}, 0, 1)
Output: []float64{1}
```

**Example 3:**

```
Input:  ClampAll([]int{}, 0, 3)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`cmp.Ordered`** | The stdlib constraint for types supporting `<`, `<=`, `>`, `>=`. |
| 2 | **Reusing a scalar rule** | Element-wise transforms are the scalar rule applied inside a loop. |
| 3 | **No aliasing** | Reused from earlier: build a new slice rather than writing into `s`. |

## Hint

Apply the `Clamp` rule per element, appending into a fresh slice.

## Validate

```bash
make verify
```

# To Floats

**Level:** junior  
**Topic:** 03-generics

## Context

A charting library only accepts `[]float64`. Counters arrive as `[]int` or `[]int64`.

## Task

Implement the stub(s) in [tofloats.go](tofloats.go):

1. Implement `ToFloats`, returning a `[]float64` with one converted element per input element.
2. Return an empty (non-nil) slice for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ToFloats([]int{1, 2})
Output: []float64{1, 2}
```

**Example 2:**

```
Input:  ToFloats([]int64{7})
Output: []float64{7}
```

**Example 3:**

```
Input:  ToFloats([]int{})
Output: []float64{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Converting a type parameter** | `float64(v)` compiles when every type in the set converts to `float64`. |
| 2 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |
| 3 | **Result type is concrete** | The output element type is `float64`, not `T` — only the input is generic. |

## Hint

`float64(v)` is legal because every type in `Integer` converts to `float64`.

## Validate

```bash
make verify
```

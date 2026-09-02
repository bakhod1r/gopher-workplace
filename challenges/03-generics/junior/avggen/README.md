# Average

**Level:** junior  
**Topic:** 03-generics

## Context

A latency dashboard averages samples that arrive as ints in one collector and as floats in another.

## Task

Implement the stub(s) in [avggen.go](avggen.go):

1. Implement `Average`, returning the mean as a `float64`.
2. Return `0` for an empty slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Average([]int{1, 2, 3})
Output: 2
```

**Example 2:**

```
Input:  Average([]float64{1, 2})
Output: 1.5
```

**Example 3:**

```
Input:  Average([]int{})
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Converting a type parameter** | `float64(v)` compiles when every type in the set converts to `float64`. |
| 2 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |
| 3 | **Integer division** | Reused from language basics: summing in `T` then dividing would truncate for ints. |

## Hint

Accumulate in `float64`, not in `T`.

## Validate

```bash
make verify
```

# Range

**Level:** junior  
**Topic:** 03-generics

## Context

Test fixtures need an index sequence in whatever integer type the surrounding code uses.

## Task

Implement the stub(s) in [rangegen.go](rangegen.go):

1. Implement `Range`, returning `0, 1, ... n-1`.
2. Return an empty (non-nil) slice when `n <= 0`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Range(3)
Output: []int{0, 1, 2}
```

**Example 2:**

```
Input:  Range(int64(2))
Output: []int64{0, 1}
```

**Example 3:**

```
Input:  Range(0)
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Converting a constant to `T`** | `T(0)` produces a typed starting value for the loop counter. |
| 2 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |
| 3 | **Loop counter of type `T`** | The counter shares the element type, so no conversion is needed inside the loop. |

## Hint

The loop counter itself is a `T`, not an `int`.

## Validate

```bash
make verify
```

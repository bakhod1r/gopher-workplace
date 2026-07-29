# Float Width & Precision

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`0.1` is not exact in binary; the rounding error per add is ~10x larger in
`float32` than `float64`. Accumulating in `float32` drifts past a 1e-9 tolerance.
The literal `0.1` is an untyped constant — the *variable's* type decides the
precision.

## Task

Fix the single line between the markers in [sums.go](sums.go) so the running sum
holds enough precision.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumTenths(10)
Output: ~1.0
```

**Example 2:**

```
Input:  SumTenths(3)
Output: ~0.3
```

**Example 3:**

```
Input:  accumulator type
Output: float64, not float32
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Untyped constant** | `0.1` adopts the accumulator's type. |
| 2 | **float32 vs float64** | Fewer mantissa bits → faster error growth. |
| 3 | **Tolerance** | 1e-9 is tighter than float32 can hold here. |

## Hint

Declare `var total float64`.

## Validate

```bash
make verify
```

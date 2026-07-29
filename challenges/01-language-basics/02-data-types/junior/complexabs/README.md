# Complex Magnitude

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Go has built-in complex numbers. `real()` and `imag()` extract the parts;
`math.Hypot` gives the magnitude without overflow.

## Task

Implement `Magnitude(c)` returning `sqrt(re^2 + im^2)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Magnitude(complex(3,4))
Output: 5
```

_Explanation:_ sqrt(9+16)=5.

**Example 2:**

```
Input:  Magnitude(complex(0,0))
Output: 0
```

_Explanation:_ Zero vector.

**Example 3:**

```
Input:  Magnitude(complex(-3,-4))
Output: 5
```

_Explanation:_ Squares drop the sign.

**Example 4:**

```
Input:  Magnitude(complex(1,0))
Output: 1
```

_Explanation:_ Pure real.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **complex128** | Built-in type; `complex(re, im)` constructs it. |
| 2 | **real / imag** | Built-in functions extract the two float64 parts. |
| 3 | **math.Hypot** | Computes sqrt(a²+b²) safely. |

## Hint

`math.Hypot(real(c), imag(c))`.

## Validate

```bash
make verify
```

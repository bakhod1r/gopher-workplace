# Float Almost-Equal

**Level:** junior
**Topic:** 01-language-basics → 02-data-types
**Estimated time:** 10 min

## Context

A pricing test kept failing at random: `total == 0.3` was false even though the
math looked right. Floating-point values are stored in binary and can't
represent most decimals exactly, so `0.1 + 0.2` is `0.30000000000000004`.
Comparisons need a tolerance — and the absolute value comes from the standard
library `math` package.

## Task

Implement `AlmostEqual` in [almostequal.go](almostequal.go) so it reports
whether `a` and `b` differ by **less than `1e-9`**. Use **`math.Abs`** from the
`math` package for the absolute difference — add `import "math"` yourself. Do
not compare floats with `==`.

Do **not** change the function signature or the tests.

## Examples

```go
AlmostEqual(0.1+0.2, 0.3) // => true   (== would be false)
AlmostEqual(1.0, 1.0)     // => true
AlmostEqual(1.0, 1.0001)  // => false
AlmostEqual(0, 1e-12)     // => true    (within tolerance)
AlmostEqual(0, 1e-6)      // => false   (outside tolerance)
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Floats are inexact** | `float64` is binary IEEE-754; most decimal fractions have no exact representation. |
| 2 | **Never `==` on floats** | Tiny representation errors make exact equality unreliable. |
| 3 | **Importing a package** | Standard-library helpers live in packages; `import "math"` makes `math.Abs` available. |
| 4 | **`math.Abs`** | Returns the absolute value as a `float64`; compare it against the tolerance. |

## Hint

Add `import "math"` at the top of the file, then
`return math.Abs(a-b) < 1e-9`. Forgetting the import gives
`undefined: math` — Go never imports packages implicitly.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

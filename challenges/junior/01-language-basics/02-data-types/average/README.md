# Average

**Level:** junior
**Topic:** 01-language-basics → 02-data-types
**Estimated time:** 15 min

## Context

A metrics endpoint reports the mean of a batch of integer samples. The naive
version divided the integer sum by the integer count and returned `2` for the
average of `1, 2, 4` — the `.333…` vanished. The fix is about *when* you cross
from integer to floating-point.

## Task

Implement `Average` in [average.go](average.go) so that it returns the arithmetic
mean of `nums` as a `float64`:

1. Sum the values and divide by the count.
2. Convert to `float64` **before** dividing so the result keeps its fractional
   part (integer division truncates).
3. An empty or nil slice returns `0`.

Do **not** change the function signature or the tests.

## Examples

```go
Average([]int{1, 2})    // => 1.5
Average([]int{1, 2, 4}) // => 2.333…
Average(nil)            // => 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Explicit conversion** | Go never converts numeric types implicitly; `float64(sum)` is required to leave integer land. |
| 2 | **Integer division truncates** | `7 / 3` with two `int` operands is `2`, not `2.333` — the remainder is dropped before any float sees it. |
| 3 | **Conversion order** | Convert the operands *before* the `/`; converting the already-truncated integer result (`float64(7/3)`) is too late. |

## Hint

The bug hides in the division, not the sum. `float64(sum) / float64(len(nums))`
divides in floating-point; `float64(sum / len(nums))` divides first as integers
and only then widens the truncated answer. Also guard the empty case so you never
divide by zero.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

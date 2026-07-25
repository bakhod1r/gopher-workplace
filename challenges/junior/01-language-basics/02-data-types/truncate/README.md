# Truncate Float

**Level:** junior
**Topic:** 01-language-basics → 02-data-types
**Estimated time:** 10 min

## Context

A billing report needs the whole-dollar part of each line amount — `9.99`
becomes `9`, `-9.99` becomes `-9`. A teammate reached for rounding and got `10`,
which threw off the ledger. The correct primitive here is a plain numeric
conversion, whose truncation behaviour is precisely defined.

## Task

Implement `WholePart` in [truncate.go](truncate.go) so that it returns the
integer part of `amount`:

1. Convert the `float64` to an `int`, dropping the fractional part.
2. Truncation is **toward zero**: `9.99 → 9` and `-9.99 → -9` (not `-10`).
3. Exact integers like `4.0` return `4`.

Do **not** change the function signature or the tests.

## Examples

```go
WholePart(9.99)  // => 9
WholePart(-9.99) // => -9
WholePart(4.0)   // => 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **`float64`→`int` conversion** | `int(x)` drops the fraction; it does not round. `int(9.99)` is `9`. |
| 2 | **Truncation toward zero** | The fraction is removed, moving the value *toward* 0: `int(-9.99)` is `-9`, not `-10`. |
| 3 | **Conversion vs rounding** | Rounding (`math.Round`) and truncation are different operations — pick the one the spec asks for. |

## Hint

This is a single explicit conversion — `int(amount)` — no `math` package needed.
Resist reaching for a rounding helper: the requirement is truncation toward zero,
which is exactly what a `float64`→`int` conversion already does.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

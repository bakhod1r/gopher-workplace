# Shadowed Tally

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants
**Estimated time:** 15 min

## Context

A billing service adds up only the *positive* line items on an invoice — refunds
and zero-value rows are skipped. A teammate's first draft looked right but always
returned `0`, and nobody could see why until they looked at the accumulator
inside the loop.

## Task

Implement `Tally` in [shadow.go](shadow.go) so that it:

1. Returns the sum of every value in `nums` that is strictly greater than `0`.
2. Ignores zero and negative values.
3. Returns `0` for a nil or empty slice.

Do **not** change the function signature or the tests.

## Examples

```go
Tally([]int{1, -2, 3}) // => 4
Tally(nil)             // => 0
Tally([]int{0, 0, 7})  // => 7
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variable scope** | A variable declared inside an `if` or `for` block lives only until that block ends. |
| 2 | **Shadowing with `:=`** | Writing `sum := sum + n` inside the loop declares a *new* `sum` that dies each iteration; the outer accumulator never changes. |
| 3 | **Assignment vs declaration** | Use `=` (or `+=`) to mutate the existing accumulator; use `:=` only to introduce a genuinely new variable. |

## Hint

The accumulator must be declared **once**, outside the loop, and *assigned* on
each qualifying iteration. If your inner statement uses `:=`, you are creating a
fresh variable that shadows the real one — its updates vanish when the block
closes.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

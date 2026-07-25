# Value Swap

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants
**Estimated time:** 5 min

## Context

A utility package needs a one-liner that exchanges two values. Go's multiple
assignment does this without a temporary variable — your job is to wire it up.

## Task

Implement `Swap` in [swap.go](swap.go) so that it:

1. Returns `a` and `b` with their values exchanged.
2. Uses multiple assignment — no temporary variable.
3. Handles equal inputs (returns them unchanged).

Do **not** change the function signature or the tests.

## Examples

```go
Swap(1, 2)  // => 2, 1
Swap(9, 9)  // => 9, 9
Swap(-3, 7) // => 7, -3
Swap(0, 5)  // => 5, 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Multiple assignment** | `a, b = b, a` evaluates the whole right side first, then assigns — no temp needed. |
| 2 | **Assignment vs no-op** | `a, b = a, b` reassigns each variable to itself: valid, but changes nothing. |
| 3 | **Tuple return** | A function may return multiple values; the caller receives them positionally. |

## Hint

Go evaluates the whole right-hand side of an assignment *before* assigning, so
`a, b = b, a` swaps in one statement — no temp. Beware `a, b = a, b`, which
reassigns each variable to itself and changes nothing.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

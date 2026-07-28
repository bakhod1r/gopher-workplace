# Retry Budget

**Level:** junior
**Topic:** 01-language-basics → 01-variables-and-constants
**Estimated time:** 5 min

## Context

The HTTP client should retry a failed request 3 times. Declare the retry
constant carefully: in Go, `'3'` in single quotes is a rune, not the number 3 —
an easy way to make the retry budget wildly wrong.

## Task

Declare the `MaxRetries` constant in [retries.go](retries.go) so that:

1. `MaxRetries == 3` (an integer, not a character).
2. `Budget()` returns `4` (first attempt + 3 retries).

Do **not** change the function signature or the tests.

## Examples

```go
MaxRetries // => 3
Budget()   // => 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Rune literals** | `'3'` is a rune (character) constant — its value is the Unicode code point `51`, not the number 3. |
| 2 | **Integer literals** | `3` (no quotes) is an untyped integer constant with value 3. |
| 3 | **Untyped constant math** | `1 + MaxRetries` inherits the constant's kind, so a stray rune literal silently poisons the arithmetic. |

## Hint

Single quotes make a *character* constant. `'3'` is the code point for the
digit, which is `51`. Drop the quotes so it is the integer `3`.

## Validate

```bash
make verify   # fmt-check + vet + test
```

Green tests + clean `vet`/`gofmt` = challenge passed.

# Integer Parity

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Parity is the classic use of the remainder operator `%` — but negative numbers
need care, because `-7 % 2` is `-1` in Go, not `1`.

## Task

Implement `Parity(n)` returning `"even"` or `"odd"`, correct for negatives.

## Examples

```go
Parity(4)  // => "even"
Parity(3)  // => "odd"
Parity(-7) // => "odd"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Remainder `%`** | `n % 2` is 0 for even numbers. |
| 2 | **Sign of `%`** | In Go the result takes the sign of the dividend. |
| 3 | **Even test** | `n%2 == 0` works for negatives; `== 1` does not. |

## Hint

Test `n%2 == 0` for even — it holds for negatives. `n%2 == 1` fails for `-7`.

## Validate

```bash
make verify
```

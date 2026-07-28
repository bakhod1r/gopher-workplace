# Parse Money to Cents

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A billing service must never store money as a float. You parse the user's
`"12.34"` straight into integer cents, exactly.

## Task

Implement `Cents(s)` → integer cents; `"7"`→700, `"3.5"`→350, `"12.34"`→1234.
Reject >2 decimals or bad format.

## Examples

```go
Cents("12.34") // => 1234, true
Cents("7")     // => 700,  true
Cents("1.234") // => 0,    false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Integer money** | Avoid float; accumulate whole cents. |
| 2 | **Optional fraction** | 0, 1, or 2 decimal digits; pad to 2. |
| 3 | **Validation** | Reject 3+ decimals, letters, empty. |

## Hint

Parse dollars before the `.`, then 0-2 fractional digits padded to two;
`dollars*100 + frac`.

## Validate

```bash
make verify
```

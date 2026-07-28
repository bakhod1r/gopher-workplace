# Manual Itoa

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Formatting is the inverse of parsing: peel digits with `%10`, prepend `'0'+d`,
divide by 10, reverse.

## Task

Implement `Format(n)` (decimal, leading `-` for negatives). No `strconv`.

## Examples

```go
Format(42)   // => "42"
Format(-17)  // => "-17"
Format(0)    // => "0"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Digit extraction** | `n%10` is the last digit. |
| 2 | **Char mapping** | `byte('0'+d)`. |
| 3 | **Sign + reverse** | Handle negatives; digits build in reverse. |

## Hint

Work with the absolute value, collect digits, reverse, prepend `-` if negative.

## Validate

```bash
make verify
```

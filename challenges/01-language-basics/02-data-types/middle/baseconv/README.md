# Base Conversion

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Writing a number in base b is repeated division: the remainders are the digits,
least significant first, so you reverse at the end.

## Task

Implement `Format(n, base)` for base 2..16, lowercase digits. `Format(0,b)="0"`.

## Examples

```go
Format(5, 2)   // => "101"
Format(255, 16)// => "ff"
Format(8, 8)   // => "10"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Repeated division** | `n%base` is the next digit, `n/=base`. |
| 2 | **Digit mapping** | 0-9 then a-f via char arithmetic. |
| 3 | **Reverse order** | Digits come out least-significant first. |

## Hint

Collect `"0123456789abcdef"[n%base]`, divide, then reverse the bytes.

## Validate

```bash
make verify
```

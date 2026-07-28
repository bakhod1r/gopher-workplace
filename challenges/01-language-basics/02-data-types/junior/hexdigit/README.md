# Hex Digit

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

Hex characters are contiguous: `'0'..'9'` then `'a'..'f'`. You reach them with
byte arithmetic on character literals.

## Task

Implement `Digit(n)` returning the lowercase hex char for `0..15`, else `'?'`.

## Examples

```go
Digit(0)  // => '0'
Digit(10) // => 'a'
Digit(15) // => 'f'
Digit(16) // => '?'
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Char arithmetic** | `'0' + n` yields the n-th digit character. |
| 2 | **byte type** | ASCII fits in a `byte` (uint8). |
| 3 | **Range guard** | Reject n outside 0..15. |

## Hint

`'0' + byte(n)` for 0–9; `'a' + byte(n-10)` for 10–15.

## Validate

```bash
make verify
```

# Base64 Digit Offset

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A base64 decoder maps characters to 6-bit values. Digits `0-9` are values
**52-61**, but the code adds 53, so every digit is off by one and decoded bytes
are corrupt.

## Task

Fix the digit offset between the markers in [base64val.go](base64val.go).

## Examples

```go
Value('0') // => 52
Value('9') // => 61
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Base64 alphabet** | A-Z=0-25, a-z=26-51, 0-9=52-61, +=62, /=63. |
| 2 | **Range offsets** | Each run starts where the previous ended. |
| 3 | **Off-by-one** | 'a' is 26 (after 25 letters), '0' is 52. |

## Hint

`return int(c-'0') + 52`.

## Validate

```bash
make verify
```

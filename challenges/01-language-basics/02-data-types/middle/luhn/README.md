# Luhn Checksum

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Payment forms pre-validate card numbers with the Luhn checksum before hitting the
gateway — cheap client-side rejection of typos.

## Task

Implement `Valid(s)`: double every second digit from the right (subtract 9 if
>9), sum all, and require the total `% 10 == 0`. Non-digits/empty → false.

## Examples

```go
Valid("4539148803436467") // => true
Valid("79927398713")      // => true
Valid("79927398710")      // => false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Right-to-left scan** | Position parity decides doubling. |
| 2 | **Digit value** | `c-'0'`, reject non-digits. |
| 3 | **Checksum mod 10** | Valid when the sum is a multiple of 10. |

## Hint

Walk from the last digit; double every second one, subtract 9 if the doubled
value exceeds 9; sum; check `sum%10==0`.

## Validate

```bash
make verify
```

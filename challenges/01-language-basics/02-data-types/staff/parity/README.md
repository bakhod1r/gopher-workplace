# Parity Bit

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A serial link appends a parity bit: 1 for an odd number of set bits. The code
**sums** the bits and returns the count, so it reports 3 instead of 1 for
`0b111` — not a parity bit at all.

## Task

Fix the accumulation between the markers in [parity.go](parity.go) to XOR.

## Examples

```go
Parity(0b111) // => 1
Parity(0xFF)  // => 0
Parity(1)     // => 1
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Parity = XOR of bits** | Odd count → 1, even → 0. |
| 2 | **XOR accumulation** | `p ^= x & 1`. |
| 3 | **Sum vs parity** | Count isn't the single parity bit. |

## Hint

`p ^= int(x & 1)`.

## Validate

```bash
make verify
```

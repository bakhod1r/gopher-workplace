# Population Count

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Counting set bits (popcount) is a bit-manipulation staple. The trick
`x &= x-1` clears the lowest set bit, so the loop runs once per 1-bit.

## Task

Implement `Count(x)` returning the number of 1-bits.

## Examples

```go
Count(0)       // => 0
Count(0b1011)  // => 3
Count(0xFF)    // => 8
Count(^uint64(0)) // => 64
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Bit AND** | `x & (x-1)` clears the lowest set bit. |
| 2 | **Loop bound** | Iterations = number of set bits. |
| 3 | **Unsigned shift** | Alternatively shift right and mask `x&1`. |

## Hint

`for x != 0 { x &= x - 1; n++ }`.

## Validate

```bash
make verify
```

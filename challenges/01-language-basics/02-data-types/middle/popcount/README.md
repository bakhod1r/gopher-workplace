# Population Count

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

Counting set bits (popcount) is a bit-manipulation staple. The trick
`x &= x-1` clears the lowest set bit, so the loop runs once per 1-bit.

## Task

Implement `Count(x)` returning the number of 1-bits.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Count(0b1011)
Output: 3
```

_Explanation:_ three set bits

**Example 2:**

```
Input:  Count(0xFF)
Output: 8
```

_Explanation:_ all low 8 bits set

**Example 3:**

```
Input:  Count(^uint64(0))
Output: 64
```

_Explanation:_ every bit set

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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

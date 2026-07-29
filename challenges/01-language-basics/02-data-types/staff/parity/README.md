# Parity Bit

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A serial link appends a parity bit: 1 for an odd number of set bits. The code
**sums** the bits and returns the count, so it reports 3 instead of 1 for
`0b111` — not a parity bit at all.

## Task

Fix the accumulation between the markers in [parity.go](parity.go) to XOR.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  1
Output: 1
```

**Example 2:**

```
Input:  3
Output: 0
```

_Explanation:_ Two set bits -> even parity.

**Example 3:**

```
Input:  0xFF
Output: 0
```

_Explanation:_ Eight set bits -> even.

**Example 4:**

```
Input:  0x80000000
Output: 1
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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

# Fletcher-16 Modulus

**Level:** senior
**Topic:** 01-language-basics → 02-data-types

## Context

A firmware image uses a Fletcher-16 checksum. The running sums are reduced mod
`256` instead of the spec's `255`, so every checksum disagrees with the
reference implementation and the bootloader rejects valid images.

## Task

Fix the two modulus operations between the markers in
[fletcher16.go](fletcher16.go) to use `255`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  "abcde"
Output: 0xC8F0
```

**Example 2:**

```
Input:  ""
Output: 0x0000
```

**Example 3:**

```
Input:  "a"
Output: 0x6161
```

_Explanation:_ sum1=sum2=97

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Fletcher sums** | Two accumulators, one of running totals. |
| 2 | **Modulus 255** | The algorithm specifies mod 255, not 256. |
| 3 | **Combine** | `(sum2 << 8) |

## Hint

Change both `% 256` to `% 255`.

## Validate

```bash
make verify
```

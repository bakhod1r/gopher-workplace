# MinInt Signed Shift

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Two's complement puts the most-negative value at `-1 << 63` for a 64-bit word.
`-1 << 62` is only a quarter of the way there. The asymmetry (`|MinInt| > MaxInt`)
is a memory-model fact the compiler encodes in the constant.

## Task

Fix the single line between the markers in [limits.go](limits.go) so `MinInt`
equals `math.MinInt64`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MinInt
Output: most negative 64-bit value
```

**Example 2:**

```
Input:  SymmetricTo()
Output: true
```

**Example 3:**

```
Input:  shift amount
Output: 63, not 62
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Signed shift** | `-1 << 63` sets only the sign bit. |
| 2 | **Two's complement asymmetry** | `MinInt` has no positive mirror. |
| 3 | **Constant width** | The value must fit int64 exactly. |

## Hint

`-1 << 63`.

## Validate

```bash
make verify
```

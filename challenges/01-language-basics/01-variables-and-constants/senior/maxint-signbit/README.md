# Max Int Sign Bit

**Level:** senior
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`int(^uint(0))` is all bits set — which as a *signed* int is `-1`, not the
maximum. The largest signed value needs the sign bit cleared with `>> 1`.

## Task

Fix the single line between the markers in [limits.go](limits.go) so `MaxInt`
is the largest positive int.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MaxInt
Output: largest positive int
```

**Example 2:**

```
Input:  Overflows()
Output: true
```

**Example 3:**

```
Input:  MaxInt+1
Output: wraps negative
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Two's complement** | All-bits-set signed = -1. |
| 2 | **Sign-bit shift** | `^uint(0) >> 1` clears the top bit. |
| 3 | **Conversion order** | Shift as `uint`, then convert to `int`. |

## Hint

`int(allBits >> 1)`.

## Validate

```bash
make verify
```

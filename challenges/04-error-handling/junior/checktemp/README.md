# Temperature Range

**Level:** junior
**Topic:** 04-error-handling

## Context

A cold-chain sensor reports readings in Celsius. Anything outside the physical range of the probe is a hardware fault.

## Task

Implement `CheckTemp` in [checktemp.go](checktemp.go):

1. Return `ErrBelowRange` when `c` is below -40.
2. Return `ErrAboveRange` when `c` is above 85.
3. Return nil for readings within `[-40, 85]`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CheckTemp(20)
Output: nil
```

**Example 2:**

```
Input:  CheckTemp(-50)
Output: ErrBelowRange
```

**Example 3:**

```
Input:  CheckTemp(100)
Output: ErrAboveRange
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Float comparisons** | Range bounds on `float64` compare directly. |
| 2 | **Inclusive bounds** | Both -40 and 85 are valid readings. |
| 3 | **Separate sentinels** | Under-range and over-range are different faults. |

## Hint

The bounds themselves are valid — strict comparisons only.

## Validate

```bash
make verify
```

# Sum Of Positives

**Level:** junior
**Topic:** 04-error-handling

## Context

A shipping calculator totals package weights. A negative weight is corrupt input and invalidates the whole shipment.

## Task

Implement `SumPositive` in [sumpositive.go](sumpositive.go):

1. Return the sum of all values and nil when every value is `>= 0`.
2. Return `0` and `ErrNegativeValue` as soon as a negative value appears.
3. Return `0` and nil for an empty or nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumPositive([]int{1, 2, 3})
Output: 6, nil
```

**Example 2:**

```
Input:  SumPositive([]int{1, -2})
Output: 0, ErrNegativeValue
```

**Example 3:**

```
Input:  SumPositive(nil)
Output: 0, nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Validate while accumulating** | One pass does both jobs. |
| 2 | **Abort semantics** | A partial sum is not returned on failure. |
| 3 | **Empty input** | An empty sum is 0, and that is not an error. |

## Hint

On failure the total is discarded — return the zero value, not what was accumulated so far.

## Validate

```bash
make verify
```

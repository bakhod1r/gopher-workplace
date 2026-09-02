# Round To Integer

**Level:** middle  
**Topic:** 03-generics

## Context

An invoice rounds line totals to whole cents, and the accounting rules say halves round away from zero.

## Task

Implement the stub(s) in [roundgen.go](roundgen.go):

1. Implement `RoundHalfUp`, keeping the caller's float type.
2. Halves round away from zero: `0.5` becomes `1`, `-0.5` becomes `-1`.
3. Use `math.Round`, which already has these semantics.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  RoundHalfUp(2.5)
Output: 3
```

**Example 2:**

```
Input:  RoundHalfUp(-2.5)
Output: -3
```

**Example 3:**

```
Input:  RoundHalfUp(2.4)
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bridging to `math`** | `math` takes `float64`, so convert out and back. |
| 2 | **`float32` round-trips** | Converting `float32` up and back down is exact for values in range. |
| 3 | **Half-away-from-zero** | `math.Round` is not banker's rounding — it rounds halves away from zero. |

## Hint

Convert to `float64`, call `math.Round`, convert back to `T`.

## Validate

```bash
make verify
```

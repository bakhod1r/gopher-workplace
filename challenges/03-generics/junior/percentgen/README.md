# Percent

**Level:** junior  
**Topic:** 03-generics

## Context

A progress bar shows completion as a percentage. Some callers count items as ints, others track byte ratios as floats.

## Task

Implement the stub(s) in [percentgen.go](percentgen.go):

1. Implement `Percent`, returning `part` as a percentage of `whole`.
2. Return `0` when `whole` is zero.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Percent(1, 4)
Output: 25
```

**Example 2:**

```
Input:  Percent(3, 3)
Output: 100
```

**Example 3:**

```
Input:  Percent(1, 0)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Converting a type parameter** | `float64(v)` compiles when every type in the set converts to `float64`. |
| 2 | **Divide-by-zero guard** | Integer division by zero panics; float division yields NaN — guard covers both. |
| 3 | **Union constraints** | `~int | ~float64` lists the types a parameter may take; only operations all of them support are allowed. |

## Hint

Convert both operands before dividing, and guard `whole == 0`.

## Validate

```bash
make verify
```

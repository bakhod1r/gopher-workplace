# Scaling A Unit

**Level:** middle  
**Topic:** 03-generics

## Context

A physics helper scales distances by a dimensionless factor. Multiplying two `Meters` together would be meaningless, so the API must not allow it.

## Task

Implement the stub(s) in [unitmulgen.go](unitmulgen.go):

1. Implement `Times`, multiplying a unit value by a plain `float64`.
2. Implement `SumUnits`, totalling values of the same unit.
3. Note that the factor is deliberately not a `T`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Times(Meters(2), 3)
Output: Meters(6)
```

**Example 2:**

```
Input:  SumUnits([]Meters{1, 2})
Output: Meters(3)
```

**Example 3:**

```
Input:  SumUnits([]Meters{})
Output: Meters(0)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Dimensional analysis in types** | Scaling by a scalar keeps the unit; multiplying two units would change it. |
| 2 | **Deliberate asymmetry** | Addition takes two `T`s; multiplication takes a `T` and a scalar. |
| 3 | **The `~` token** | `~int` admits every type whose underlying type is `int`. |

## Hint

The factor is a `float64` on purpose — `Times(Meters, Meters)` should not typecheck.

## Validate

```bash
make verify
```

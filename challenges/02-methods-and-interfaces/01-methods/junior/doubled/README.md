# Doubled

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

An animation engine doubles frame durations when in slow-motion mode. The
duration is stored as a custom float type.

## Task

Implement `Double` on `MyFloat` in [doubled.go](doubled.go):

1. Return the value multiplied by 2.
2. Value receiver — do **not** mutate the original.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MyFloat(3.5).Double()
Output: 7.0
```

**Example 2:**

```
Input:  MyFloat(0).Double()
Output: 0
```

**Example 3:**

```
Input:  MyFloat(-2.5).Double()
Output: -5.0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Value receiver** | Returns a new value; original unchanged. |
| 2 | **Defined type arithmetic** | `MyFloat * 2` works because `MyFloat`'s underlying type is `float64`. |

## Hint

`return f * 2` — Go allows arithmetic on defined numeric types.

## Validate

```bash
make verify
```

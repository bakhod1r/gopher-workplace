# Celsius to Fahrenheit

**Level:** junior
**Topic:** 01-language-basics → 02-data-types

## Context

`c*9/5` in integers truncates. Converting to `float64` first keeps the fraction,
so `37°C` gives `98.6`, not `98`.

## Task

Implement `ToF(c)` returning `c*9/5 + 32` as a `float64`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ToF(0)
Output: 32
```

_Explanation:_ Freezing point.

**Example 2:**

```
Input:  ToF(100)
Output: 212
```

_Explanation:_ Boiling point.

**Example 3:**

```
Input:  ToF(37)
Output: 98.6
```

_Explanation:_ float64(37)*9/5+32 keeps the .6; integer math would give 98.

**Example 4:**

```
Input:  ToF(-40)
Output: -40
```

_Explanation:_ The scales cross at -40.

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Type conversion** | `float64(c)` widens the int before dividing. |
| 2 | **Integer vs float division** | `9/5` in ints is 1; in floats 1.8. |
| 3 | **Mixed arithmetic** | Go has no implicit int↔float; convert explicitly. |

## Hint

`return float64(c)*9/5 + 32`.

## Validate

```bash
make verify
```

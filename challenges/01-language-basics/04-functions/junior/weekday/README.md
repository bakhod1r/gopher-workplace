# Weekday Name

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _control-flow_

## Context

A `switch` on a value maps each case to a result, with `default` handling everything out of range.

## Task

Implement `Weekday` in [weekday.go](weekday.go) mapping 1–7 to day names.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Weekday(1)
Output: "Monday"
```

**Example 2:**

```
Input:  Weekday(7)
Output: "Sunday"
```

**Example 3:**

```
Input:  Weekday(8)
Output: "Unknown"
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **expression switch** | `switch d { case 1: }`. |
| 2 | **default case** | Out-of-range → "Unknown". |
| 3 | **value mapping** | Each day number to a name. |

## Hint

`switch d { case 1: return "Monday"; ...; default: return "Unknown" }`.

## Validate

```bash
make verify
```

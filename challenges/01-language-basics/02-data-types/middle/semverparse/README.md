# Semver Parse

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

A deploy tool reads image tags like `1.4.10` and needs the three numeric parts
to compare releases. You parse the string into integers, rejecting malformed
tags.

## Task

Implement `Parse(s)` returning `(major, minor, patch, ok)`; `ok=false` unless the
string is exactly three dot-separated non-negative integers.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Parse("1.4.10")
Output: (1, 4, 10, true)
```

_Explanation:_ three integer parts

**Example 2:**

```
Input:  Parse("2.0")
Output: (0, 0, 0, false)
```

_Explanation:_ only two parts

**Example 3:**

```
Input:  Parse("1.x.0")
Output: (0, 0, 0, false)
```

_Explanation:_ non-digit component

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Byte scanning** | Walk the string, accumulating digits per field. |
| 2 | **Digit fold** | `n = n*10 + (c-'0')`. |
| 3 | **Strict validation** | Exactly two dots, digits only, no empty field. |

## Hint

Track a current number and a field count; on `.` advance, on a digit fold, else
fail. Require exactly 3 fields, each non-empty.

## Validate

```bash
make verify
```

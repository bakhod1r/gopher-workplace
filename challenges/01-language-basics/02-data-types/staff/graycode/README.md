# Gray Code Encoding

**Level:** staff
**Topic:** 01-language-basics → 02-data-types

## Context

A rotary encoder outputs Gray code so only one bit changes per step. The formula
is `x ^ (x >> 1)`, but the code shifts **left**, breaking the single-bit-change
property.

## Task

Fix the shift between the markers in [graycode.go](graycode.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  2
Output: 3
```

**Example 2:**

```
Input:  3
Output: 2
```

**Example 3:**

```
Input:  4
Output: 6
```

**Example 4:**

```
Input:  7
Output: 4
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Gray code** | `g = x ^ (x >> 1)`. |
| 2 | **Single-bit change** | Consecutive codes differ by one bit. |
| 3 | **Shift direction** | Right shift folds high bits down. |

## Hint

`x ^ (x >> 1)`.

## Validate

```bash
make verify
```

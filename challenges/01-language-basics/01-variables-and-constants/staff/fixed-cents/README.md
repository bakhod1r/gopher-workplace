# Fixed-Point Conversion Order

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`int64(dollars) * 100` truncates the dollars to a whole number *before* scaling,
so `2.50` becomes 200 cents. Scale in floating point, then convert.

## Task

Fix the single line between the markers in [money.go](money.go) so fractional
dollars survive.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Cents(1.99)
Output: 199
```

**Example 2:**

```
Input:  Cents(0.5)
Output: 50
```

_Explanation:_ Truncating dollars first would give 0.

**Example 3:**

```
Input:  Cents(2.0)
Output: 200
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Conversion truncates** | `int64(2.50)` is 2, dropping the cents. |
| 2 | **Order of operations** | Multiply in float, convert last. |
| 3 | **Fixed-point money** | Represent currency as integer minor units. |

## Hint

`int64(dollars * 100)`.

## Validate

```bash
make verify
```

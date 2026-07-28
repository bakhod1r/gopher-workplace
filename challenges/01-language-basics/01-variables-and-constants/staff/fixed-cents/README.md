# Fixed-Point Conversion Order

**Level:** staff
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

`int64(dollars) * 100` truncates the dollars to a whole number *before* scaling,
so `2.50` becomes 200 cents. Scale in floating point, then convert.

## Task

Fix the single line between the markers in [money.go](money.go) so fractional
dollars survive.

## Examples

```go
Cents(1.00)  // => 100
Cents(2.50)  // => 250
Cents(0.99)  // => 99
```

## Topics to Master

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

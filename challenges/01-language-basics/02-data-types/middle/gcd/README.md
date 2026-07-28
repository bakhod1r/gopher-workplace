# Greatest Common Divisor

**Level:** middle
**Topic:** 01-language-basics → 02-data-types

## Context

The Euclidean algorithm repeatedly replaces the pair `(a, b)` with `(b, a%b)`
until `b` is 0. Signs must be normalized so the result is non-negative.

## Task

Implement `GCD(a, b)` (non-negative, `GCD(0,0)=0`).

## Examples

```go
GCD(12, 8) // => 4
GCD(17, 5) // => 1
GCD(0, 9)  // => 9
GCD(-12, 8)// => 4
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Modulo loop** | `a, b = b, a%b` until b==0. |
| 2 | **Sign handling** | Take absolute values first. |
| 3 | **Multiple assignment** | Update both in one statement. |

## Hint

Loop `for b != 0 { a, b = b, a%b }`; return `abs(a)`.

## Validate

```bash
make verify
```

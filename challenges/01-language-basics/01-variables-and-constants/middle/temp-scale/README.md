# Typed Temperature Constants

**Level:** middle
**Topic:** 01-language-basics → 01-variables-and-constants

## Context

Typing a constant (`Celsius`) makes it participate in a domain type's API and
prevents accidental mixing with plain floats.

## Task

In [temp.go](temp.go):

1. Define `AbsoluteZero Celsius = -273.15` and `Boiling Celsius = 100`.
2. Implement `Valid(c)` returning true when `c >= AbsoluteZero`.

## Examples

```go
AbsoluteZero      // => -273.15
Valid(-300)       // => false
Valid(0)          // => true
Valid(Boiling)    // => true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Typed constants** | `Celsius = -273.15` is a `Celsius`, not a bare float. |
| 2 | **Domain types** | Prevents mixing Celsius with raw numbers. |
| 3 | **Comparison** | Typed constants compare within their own type. |

## Hint

Put the type on each constant: `AbsoluteZero Celsius = -273.15`. Then `Valid`
is a single `>=` comparison.

## Validate

```bash
make verify
```

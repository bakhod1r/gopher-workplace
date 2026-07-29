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

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Valid(-273.15)
Output: true
```

**Example 2:**

```
Input:  Valid(-300)
Output: false
```

**Example 3:**

```
Input:  Valid(100)
Output: true
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

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

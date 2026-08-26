# Celsius

**Level:** junior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A weather API returns temperatures in Celsius. The display layer needs
Fahrenheit. Instead of a standalone function, attach the conversion directly to
the `Celsius` type.

## Task

Implement `ToFahrenheit` on `Celsius` in [celsius.go](celsius.go):

1. Return `C × 9/5 + 32`.
2. `Celsius` is a **defined type** (`type Celsius float64`).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Celsius(0).ToFahrenheit()
Output: 32
```

**Example 2:**

```
Input:  Celsius(100).ToFahrenheit()
Output: 212
```

**Example 3:**

```
Input:  Celsius(-40).ToFahrenheit()
Output: -40
```

_Explanation:_ −40 is where Celsius and Fahrenheit meet.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Methods on defined types** | `type Celsius float64` gets its own methods. |
| 2 | **Value receiver** | Read-only computation — no mutation. |
| 3 | **Numeric conversion** | `float64(c)` converts the defined type to its underlying type. |

## Hint

`return float64(c)*9/5 + 32` — convert the receiver to `float64` for arithmetic.

## Validate

```bash
make verify
```

# Variadic Sum

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

A variadic parameter `nums ...int` collects any number of trailing arguments
into a slice. A caller can also spread an existing slice with `xs...`.

## Task

Implement `Sum` in [sumv.go](sumv.go) so it adds up every argument.

Do **not** change the function signature or the tests.

## Examples

```go
Sum()          // => 0
Sum(1, 2, 3)   // => 6
Sum(xs...)     // spread a slice
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic parameter** | `nums ...int` is a `[]int` inside the function. |
| 2 | **Zero arguments** | Calling with none yields an empty slice, so the sum is 0. |
| 3 | **Spread operator** | `Sum(xs...)` passes a slice's elements individually. |

## Hint

Range over `nums`, accumulating into a total that starts at 0.

## Validate

```bash
make verify
```

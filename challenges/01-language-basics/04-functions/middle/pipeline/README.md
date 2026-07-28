# Function Pipeline

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _variadic_

## Context

A variadic of functions lets a caller build a processing pipeline; the value
flows through each stage in order.

## Task

Implement `Pipe` in [pipeline.go](pipeline.go).

Do **not** change the function signature or the tests.

## Examples

```go
Pipe(3, inc, dbl) // => 8
Pipe(5)           // => 5
Pipe(2, dbl, dbl) // => 8
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Variadic of functions** | `fns ...func(int) int`. |
| 2 | **Left-to-right application** | Thread `x` through each. |
| 3 | **Identity on empty** | No stages returns x unchanged. |

## Hint

`for _, f := range fns { x = f(x) }; return x`.

## Validate

```bash
make verify
```

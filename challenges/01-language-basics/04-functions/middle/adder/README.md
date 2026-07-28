# Adder Factory

**Level:** middle
**Topic:** 01-language-basics → 04-functions · _closures_

## Context

A closure can capture a parameter of its factory, freezing it into the returned
function — partial application.

## Task

Implement `Adder` in [adder.go](adder.go).

Do **not** change the function signature or the tests.

## Examples

```go
add10 := Adder(10); add10(5) // => 15
Adder(1)(5)                  // => 6
Adder(0)(9)                  // => 9
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Capture a parameter** | `base` is remembered by the closure. |
| 2 | **Partial application** | Bind one argument now, the rest later. |
| 3 | **Returned func type** | `func(int) int`. |

## Hint

`return func(x int) int { return base + x }`.

## Validate

```bash
make verify
```

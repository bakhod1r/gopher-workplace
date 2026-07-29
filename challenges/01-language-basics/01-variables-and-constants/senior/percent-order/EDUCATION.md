# Evaluation order with integer division

## Intuition

`*` and `/` share the same precedence and associate **left to right**. With
integers, doing the division first can truncate to 0 before you ever multiply:

```go
part, total := 1, 4
part / total * 100 // (1/4)*100 == 0*100 == 0
part * 100 / total // (1*100)/4 == 25   ✓
```

Same operators, same operands — only the order differs, and integer truncation
makes that order decisive.

## Approach

1. Integer `part / total` truncates to 0 before the `* 100`.
2. Multiply first: `part * 100 / total`.

## Solution

```go
func Percent(part, total int) int {
	return part * 100 / total
}
```

## Walkthrough

`Percent(1, 4)`: the bug computes `1/4 = 0`, then `*100 = 0`. Multiplying first gives `100/4 = 25`.

## Pitfalls

- Watch for intermediate overflow the other way: `part * 100` must not overflow
  the type. Widen if needed.
- Floating point avoids truncation but adds rounding — pick per use.
- Parenthesize when intent matters: `(part * 100) / total`.

# Evaluation order with integer division

## The idea

`*` and `/` share the same precedence and associate **left to right**. With
integers, doing the division first can truncate to 0 before you ever multiply:

```go
part, total := 1, 4
part / total * 100 // (1/4)*100 == 0*100 == 0
part * 100 / total // (1*100)/4 == 25   ✓
```

Same operators, same operands — only the order differs, and integer truncation
makes that order decisive.

## Why it matters

Percentages, ratios, and scaled averages computed in integers all hit this. The
"multiply before divide" rule preserves the numerator's magnitude so the divide
has something to work with.

## Watch out

- Watch for intermediate overflow the other way: `part * 100` must not overflow
  the type. Widen if needed.
- Floating point avoids truncation but adds rounding — pick per use.
- Parenthesize when intent matters: `(part * 100) / total`.

## Try it yourself

```go
7 / 2 * 2   // 6  (3*2)
7 * 2 / 2   // 7
1 * 100 / 3 // 33
```

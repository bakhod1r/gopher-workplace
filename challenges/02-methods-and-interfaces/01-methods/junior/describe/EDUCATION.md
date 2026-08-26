# Same Method Name, Different Types

## Intuition

In Go, different types can have methods with the same name. `Circle.Describe()`
and `Square.Describe()` coexist without conflict because methods are scoped to
their receiver type. This is not overloading (Go doesn't have that) — it's just
separate method sets on separate types.

## Approach

1. Use `fmt.Sprintf` with `%g` for clean float formatting.
2. Each type returns its own label.

## Solution

```go
func (c Circle) Describe() string {
	return fmt.Sprintf("Circle(radius=%g)", c.Radius)
}

func (s Square) Describe() string {
	return fmt.Sprintf("Square(side=%g)", s.Side)
}
```

## Walkthrough

For `Circle{5}`:
- `%g` formats `5.0` as `"5"`.
- Result: `"Circle(radius=5)"`.

## Pitfalls

- Using `%f` gives `"5.000000"` — trailing zeros break the test.
- Using `%v` gives `"5"` for integers but may differ for edge cases — `%g` is
  the reliable choice.
- Methods don't "override" each other — Go has no inheritance. They simply live
  on separate types.

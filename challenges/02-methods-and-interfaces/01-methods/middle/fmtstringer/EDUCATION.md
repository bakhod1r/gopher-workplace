# Implicit Interfaces

## Intuition

Go interfaces are satisfied implicitly. By writing `String() string`, your type
automatically implements `fmt.Stringer`. `fmt.Print` uses reflection to check
for this method and calls it if present.

## Approach

1. Format and return.

## Solution

```go
func (p Point) String() string {
	return fmt.Sprintf("(%d,%d)", p.X, p.Y)
}
```

## Walkthrough

- `fmt.Sprintf("(%d,%d)", 1, 2)` → `"(1,2)"`.

## Pitfalls

- Calling `fmt.Sprint(p)` inside `String()` causes infinite recursion, as
  `Sprint` will call `String()` again.

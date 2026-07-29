# Typed constants

## Intuition

Putting a type on a constant makes it a value of a *domain* type, not a bare
number:

```go
type Celsius float64
const (
	AbsoluteZero Celsius = -273.15
	Boiling      Celsius = 100
)
```

`AbsoluteZero` is a `Celsius`, so it slots naturally into any API that speaks
`Celsius` and resists being mixed up with, say, a `Fahrenheit` or a raw float.

## Approach

1. Declare `AbsoluteZero` and `Boiling` as typed `Celsius` constants.
2. `Valid` checks `c >= AbsoluteZero`.

## Solution

```go
type Celsius float64

const (
	AbsoluteZero Celsius = -273.15
	Boiling      Celsius = 100
)

func Valid(c Celsius) bool {
	return c >= AbsoluteZero
}
```

## Walkthrough

`Valid(-300)` is below absolute zero → false; exactly -273.15 is valid.

## Pitfalls

- Give each constant the type, or type the first and let bare lines inherit it.
- Typed and untyped constants mix in arithmetic, but two **different** named
  types do not: `Celsius + Fahrenheit` will not compile.
- Comparisons (`>=`, `<`) work within the type.

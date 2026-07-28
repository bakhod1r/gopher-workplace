# Typed constants

## The idea

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

## Why it matters

Domain types turn unit mistakes into compile errors:

```go
func Freeze(c Celsius) bool { return c <= 0 }
Freeze(AbsoluteZero) // fine
Freeze(98.6)         // 98.6 is an untyped constant, converts to Celsius — careful!
```

Untyped literals still convert implicitly, but a `Fahrenheit` variable would be
rejected — the type wall does its job for *variables*, the common source of bugs.

## Watch out

- Give each constant the type, or type the first and let bare lines inherit it.
- Typed and untyped constants mix in arithmetic, but two **different** named
  types do not: `Celsius + Fahrenheit` will not compile.
- Comparisons (`>=`, `<`) work within the type.

## Try it yourself

```go
type Meters float64
const Marathon Meters = 42195
```

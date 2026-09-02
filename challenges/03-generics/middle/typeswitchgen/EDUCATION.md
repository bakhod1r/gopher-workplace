# Type Switching On A Parameter

## Intuition

Type parameters are static; type switches are dynamic. Converting to `any` moves the value into the dynamic world, at the cost of the boxing this style of code was meant to avoid.

## Approach

1. Convert `v` to `any`.
2. Switch on its dynamic type.
3. Return the matching label.

## Solution

```go
func Describe[T any](v T) string {
	switch any(v).(type) {
	case int, int64:
		return "integer"
	case float32, float64:
		return "float"
	case string:
		return "string"
	default:
		return "other"
	}
}
```

## Walkthrough

`Describe(1)` boxes the `int` and matches the first case; a `Celsius` would fall through to `"other"`.

## Pitfalls

- Writing `switch v.(type)` on a type parameter, which does not compile.
- Expecting named types to match their underlying type in a switch.
- Reaching for this pattern when a constraint would express the intent statically.

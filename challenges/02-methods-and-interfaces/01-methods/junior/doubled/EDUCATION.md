# Value Receiver Returns

## Intuition

A value receiver method that returns a new value (instead of mutating) follows a
functional style. The caller decides what to do with the result:
`result := f.Double()`.

## Approach

1. Multiply `f` by 2.
2. Return the result.

## Solution

```go
func (f MyFloat) Double() MyFloat {
	return f * 2
}
```

## Walkthrough

For `MyFloat(3.5)`:
- `3.5 * 2` = `7.0`.

## Pitfalls

- You might think you need `float64(f) * 2` — but Go allows arithmetic on
  defined numeric types directly.
- Returning `float64` instead of `MyFloat` would be a type mismatch.

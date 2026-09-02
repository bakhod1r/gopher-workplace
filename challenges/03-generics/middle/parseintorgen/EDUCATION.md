# Parse Or Default

## Intuition

The default argument doubles as the inference source, which is why the explicit form is only needed when the default is an untyped constant.

## Approach

1. Parse with `strconv.ParseInt` in base 10.
2. Return `def` on error.
3. Otherwise convert and return.

## Solution

```go
func ParseOr[T Integer](s string, def T) T {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return def
	}
	return T(n)
}
```

## Walkthrough

`ParseOr("5", Retries(1))` infers `T = Retries` from the default and returns `Retries(5)`.

## Pitfalls

- Ignoring the parse error and returning a zero value.
- Returning `int64` and forcing callers to convert.
- Assuming the conversion is range-checked — it is not.

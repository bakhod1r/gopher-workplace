# Strings And Bytes

## Intuition

Mixed-kind type sets are the sharpest illustration of the operation rule: the intersection of what `string` and `[]byte` support is small, and everything outside it is a compile error.

## Approach

1. `HasPrefix`: convert to `string` and delegate to the `strings` package.
2. `Size`: return `len(v)`.

## Solution

```go
func HasPrefix[T ~string | ~[]byte](v T, prefix string) bool {
	return strings.HasPrefix(string(v), prefix)
}

func Size[T ~string | ~[]byte](v T) int {
	return len(v)
}
```

## Walkthrough

`HasPrefix([]byte("POST"), "GET")` copies the bytes into a string and compares, returning false.

## Pitfalls

- Trying to `append` to `v`, which is undefined for `string`.
- Assuming the `[]byte` conversion is free in a hot loop.
- Writing two separate functions when one covers both.

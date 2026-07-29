# Mixing fixed and variadic parameters

## Intuition

Only the last parameter may be variadic; fixed parameters bind positionally before it.

## Approach

1. Add the separator before every part except the first.
2. Concatenate into the result.

## Solution

```go
func Join(sep string, parts ...string) string {
	result := ""
	for i, p := range parts {
		if i > 0 {
			result += sep
		}
		result += p
	}
	return result
}
```

## Walkthrough

`Join(",", "a","b","c")` starts with "a", then ",b", then ",c" → "a,b,c".

## Pitfalls

- Prepending `sep` unconditionally leaks a leading separator.
- With zero parts the loop never runs and `""` is correct.

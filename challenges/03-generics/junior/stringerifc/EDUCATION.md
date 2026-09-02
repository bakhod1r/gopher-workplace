# Generics Versus Interfaces

## Intuition

The practical difference is at the call site: passing `[]Tag` to `Describe` costs nothing, while `DescribeAny` forces the caller to build a new `[]fmt.Stringer` element by element.

## Approach

1. Allocate the result with capacity `len(items)`.
2. Append `it.String()` for each element.
3. Return the result.

## Solution

```go
func Describe[T fmt.Stringer](items []T) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		out = append(out, it.String())
	}
	return out
}

func DescribeAny(items []fmt.Stringer) []string {
	out := make([]string, 0, len(items))
	for _, it := range items {
		out = append(out, it.String())
	}
	return out
}
```

## Walkthrough

`Describe([]Tag{"a"})` instantiates `T = Tag` and calls the concrete method directly — no interface value is ever built.

## Pitfalls

- Assuming `[]Tag` can be passed where `[]fmt.Stringer` is wanted.
- Using a type parameter for a single value, where a plain interface is simpler.
- Constraining to `any` and type-asserting to `fmt.Stringer` inside.

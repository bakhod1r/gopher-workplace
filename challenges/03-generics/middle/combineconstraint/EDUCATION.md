# Combining Type Sets

## Intuition

Composing type sets keeps the definitions honest: adding a type to `Number` extends every constraint built from it, rather than drifting out of sync.

## Approach

1. Allocate the result with capacity `len(s)`.
2. Append `fmt.Sprint(v)` for each element.
3. Return the slice.

## Solution

```go
func Render[T NumOrText](s []T) []string {
	out := make([]string, 0, len(s))
	for _, v := range s {
		out = append(out, fmt.Sprint(v))
	}
	return out
}
```

## Walkthrough

`Render([]int{1, 2})` instantiates `T = int`, which is inside `Number` and therefore inside the union.

## Pitfalls

- Re-listing every type instead of composing the existing constraints.
- Constraining to `any`, which would accept types the renderer cannot handle.
- Assuming `fmt.Sprint` is free — it allocates and reflects.

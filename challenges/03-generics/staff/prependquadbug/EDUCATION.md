# Reverse That Costs A Square

## Intuition

Building a reversed slice by repeatedly prepending copies the whole accumulated prefix each time. The total work is `1 + 2 + ... + n`, so doubling the input quadruples the runtime.

## Approach

1. Allocate the result at full length up front.
2. Walk the input once, writing each element to the mirrored index.

## Solution

```go
func Reversed[T any](s []T) []T {
	out := make([]T, len(s))
	for i, v := range s {
		out[len(s)-1-i] = v
	}
	return out
}
```

## Walkthrough

For `n = 120000`, prepending copies about seven billion elements and allocates 120000 separate arrays; writing to `out[n-1-i]` copies 120000 elements once.

## Pitfalls

- Reversing in place and returning the caller's slice, which mutates the input.
- Appending then reversing with a second pass — correct, but two passes where one suffices.

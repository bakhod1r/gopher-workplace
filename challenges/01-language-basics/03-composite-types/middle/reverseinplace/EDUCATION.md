# In-place reversal

## Intuition

Swap symmetric pairs from the outside in, stopping at the middle:

```go
for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
	xs[i], xs[j] = xs[j], xs[i]
}
```

## Approach

1. Two indices i=0, j=len-1.
2. Swap xs[i],xs[j], move inward until i>=j.
3. Mutates the input in place.
4. Return xs.

## Solution

```go
func Reverse(xs []int) []int {
	for i, j := 0, len(xs)-1; i < j; i, j = i+1, j-1 {
		xs[i], xs[j] = xs[j], xs[i]
	}
	return xs
}
```

## Walkthrough

[1,2,3,4]: swap(0,3)->[4,2,3,1]; swap(1,2)->[4,3,2,1]; i>=j stop.

## Pitfalls

- This **mutates** the caller's slice (shared backing array); copy first if the
  original must survive.
- Multiple assignment swaps without a temp.
- `slices.Reverse` (Go 1.21+) does this generically.

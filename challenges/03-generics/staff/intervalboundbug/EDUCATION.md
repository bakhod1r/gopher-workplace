# The Range That Overlaps Its Neighbour

## Intuition

With a half-open range, `hi` is the first position *outside*. An inclusive loop claims one extra position, so ranges that should tile perfectly overlap at every seam and the total is inflated by exactly the number of ranges.

## Approach

1. Create the map on first use.
2. Walk from `lo` while the index is strictly less than `hi`.
3. Add `v` at each visited position.

## Solution

```go
func (a *Accum[T]) Add(lo, hi int, v T) {
	if a.m == nil {
		a.m = make(map[int]T)
	}
	for i := lo; i < hi; i++ {
		a.m[i] += v
	}
}

func (a *Accum[T]) At(x int) T {
	return a.m[x]
}

func (a *Accum[T]) Total() T {
	var sum T
	for _, v := range a.m {
		sum += v
	}
	return sum
}

func (a *Accum[T]) Touched() int {
	return len(a.m)
}
```

## Walkthrough

`Add(0,3,1)` and `Add(3,6,1)` should tile positions 0..5 with 1 each. The inclusive loop gives position 3 a value of 2 and also touches position 6.

## Pitfalls

- Compensating with `hi-1` at the call site instead of fixing the loop.
- Assuming a reversed range should count backwards; it must add nothing.
- Testing only with ranges separated by a gap, where the overlap cannot appear.

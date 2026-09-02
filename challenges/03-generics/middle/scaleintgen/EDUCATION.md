# Scale To A Ceiling

## Intuition

Integer scaling is a precision-versus-overflow trade: multiplying first keeps small values visible, at the cost of a wider intermediate.

## Approach

1. Return empty for an empty input.
2. Find the largest element.
3. Return a copy unchanged when it is not positive.
4. Otherwise map each element with `v*top/peak`.

## Solution

```go
func Scale[T Integer](s []T, top T) []T {
	out := make([]T, 0, len(s))
	if len(s) == 0 {
		return out
	}
	peak := s[0]
	for _, v := range s[1:] {
		if v > peak {
			peak = v
		}
	}
	if peak <= 0 {
		out = append(out, s...)
		return out
	}
	for _, v := range s {
		out = append(out, v*top/peak)
	}
	return out
}
```

## Walkthrough

`Scale([]int{1, 2, 4}, 100)` computes `1*100/4 = 25`, `2*100/4 = 50`, `4*100/4 = 100`.

## Pitfalls

- Writing `v/peak*top`, which yields all zeros for `v < peak`.
- Dividing by a zero peak.
- Scaling against the sum instead of the maximum.

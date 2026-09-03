# Throwing Away The First Rounds

## Intuition

The measurement you want is of the steady state. The first samples describe a different program — one that is still setting itself up.

## Approach

1. Clamp `n` into range.
2. Copy the remainder so the caller's array stays private.
3. Average the copy.

## Solution

```go
func Drop(samples []float64, n int) []float64 {
	if n < 0 {
		n = 0
	}
	if n > len(samples) {
		n = len(samples)
	}
	out := slices.Clone(samples[n:])
	if out == nil {
		out = []float64{}
	}
	return out
}

func StableMean(samples []float64, n int) float64 {
	rest := Drop(samples, n)
	if len(rest) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range rest {
		sum += v
	}
	return sum / float64(len(rest))
}
```

## Walkthrough

Without the clone, `Drop` would hand back a window onto the caller's array and a later write to the result would silently edit their samples.

## Pitfalls

- Slicing with an unclamped `n`, which panics past the end.
- Returning `samples[n:]` directly and aliasing the input.
- Dropping so many samples that the remaining set is too small to mean anything.

# ns/op And "Is It Actually Faster?"

## Intuition

The tool truncates ns/op, and a percentage improvement is always measured against the baseline, never against the candidate.

## Approach

1. Guard both inputs, then divide.
2. Compare the candidate against `base` scaled down by the required fraction.

## Solution

```go
func NsPerOp(elapsedNS int64, iters int) int64 {
	if iters <= 0 || elapsedNS < 0 {
		return 0
	}
	return elapsedNS / int64(iters)
}

func Faster(base, candidate int64, pct float64) bool {
	if base <= 0 {
		return false
	}
	return float64(candidate) <= float64(base)*(1-pct/100)
}
```

## Walkthrough

`Faster(100, 80, 20)` compares `80 <= 100*0.8 = 80`, true at the boundary; `Faster(100, 81, 20)` compares `81 <= 80`, false.

## Pitfalls

- Dividing by `base` and comparing to `pct` directly, which inverts the direction.
- A strict `<`, which rejects a change that hits the target exactly.
- Measuring improvement against the candidate, inflating every result.

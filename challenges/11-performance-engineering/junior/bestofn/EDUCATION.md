# Report The Best Run, Measure The Noise

## Intuition

Take the same measurement several times and the distribution has a hard floor and a long upper tail. The floor is the signal; the tail is the machine.

## Approach

1. Guard the empty slice.
2. Walk once, tracking the smallest value with its index and the largest value.
3. Divide for the spread, guarding a zero minimum.

## Solution

```go
func Best(samples []float64) (float64, int) {
	if len(samples) == 0 {
		return 0, -1
	}
	bestV, bestI := samples[0], 0
	for i, v := range samples {
		if v < bestV {
			bestV, bestI = v, i
		}
	}
	return bestV, bestI
}

func Spread(samples []float64) float64 {
	minV, _ := Best(samples)
	if len(samples) == 0 || minV == 0 {
		return 0
	}
	maxV := samples[0]
	for _, v := range samples {
		if v > maxV {
			maxV = v
		}
	}
	return maxV / minV
}
```

## Walkthrough

A strict `<` in the comparison is what keeps the earliest index on a tie: a later equal value never replaces the one already stored.

## Pitfalls

- Initialising the minimum to `0`, which no positive sample ever beats.
- Using `<=` and reporting the last tied index instead of the first.
- Reporting the minimum without the spread, hiding that the machine was too noisy to trust.

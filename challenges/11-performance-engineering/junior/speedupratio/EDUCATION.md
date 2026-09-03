# "Twice As Fast" Means One Thing

## Intuition

The speedup is how many of the new operation fit in the time of the old one. The percent change is how much of the old time you removed.

## Approach

1. Guard each formula's denominators.
2. Divide.

## Solution

```go
func Speedup(baseNS, candidateNS float64) float64 {
	if baseNS <= 0 || candidateNS <= 0 {
		return 0
	}
	return baseNS / candidateNS
}

func PercentChange(baseNS, candidateNS float64) float64 {
	if baseNS <= 0 {
		return 0
	}
	return (candidateNS - baseNS) / baseNS * 100
}
```

## Walkthrough

`PercentChange` needs no guard on the candidate: a candidate of `0` legitimately means "−100%", the whole cost removed.

## Pitfalls

- Dividing by the candidate in `PercentChange`, which reports −20% as +25%.
- Reporting a 2x win as "200% faster".
- Guarding the candidate in `PercentChange` and losing the meaningful −100% case.

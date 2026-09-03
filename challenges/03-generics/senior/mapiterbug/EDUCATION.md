# A Ranking That Changes Between Runs

## Intuition

Ranging the counts map visits keys in a randomised order, so among equally frequent values whichever comes first that run wins.

## Approach

1. Tally every element.
2. Walk `s` in order, keeping the first element with a strictly higher count.

## Solution

```go
func Mode[T comparable](s []T) (T, bool) {
	if len(s) == 0 {
		var zero T
		return zero, false
	}
	counts := make(map[T]int, len(s))
	for _, v := range s {
		counts[v]++
	}
	best, bestN := s[0], counts[s[0]]
	for _, v := range s[1:] {
		if counts[v] > bestN {
			best, bestN = v, counts[v]
		}
	}
	return best, true
}
```

## Walkthrough

For `[1 1 2 2]` both counts are 2; walking the slice keeps `1`, while ranging the map returns either.

## Pitfalls

- Ranking by ranging a map.
- Sorting the keys to fix the order — deterministic, but it changes the documented tie rule.
- Blaming flaky tests instead of the non-determinism they found.

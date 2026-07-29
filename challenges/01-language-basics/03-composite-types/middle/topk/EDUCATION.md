# Rank by frequency

## Intuition

Count with a map, then sort the distinct keys by a compound key — count
descending, then the value ascending for stable tie-breaks:

```go
sort.Slice(keys, func(i, j int) bool {
	if m[keys[i]] != m[keys[j]] { return m[keys[i]] > m[keys[j]] }
	return keys[i] < keys[j]
})
```

## Approach

1. Count word frequencies into a map.
2. Collect distinct words.
3. Sort by count descending, then word ascending on ties.
4. Clamp k to len(words).
5. Return first k.

## Solution

```go
import "sort"

func TopK(xs []string, k int) []string {
	counts := map[string]int{}
	for _, w := range xs {
		counts[w]++
	}
	words := make([]string, 0, len(counts))
	for w := range counts {
		words = append(words, w)
	}
	sort.Slice(words, func(i, j int) bool {
		if counts[words[i]] != counts[words[j]] {
			return counts[words[i]] > counts[words[j]]
		}
		return words[i] < words[j]
	})
	if k > len(words) {
		k = len(words)
	}
	if k < 0 {
		k = 0
	}
	return words[:k]
}
```

## Walkthrough

counts a:3,b:2,c:1,d:1. sort -> [a,b,c,d]. k=2 -> [a,b]; k=4 ties c<d alphabetically -> [a,b,c,d].

## Pitfalls

- Map iteration is random; you must sort for a stable ranking.
- Break ties explicitly, or equal counts come out in arbitrary order.
- For large data, a heap of size k beats a full sort.

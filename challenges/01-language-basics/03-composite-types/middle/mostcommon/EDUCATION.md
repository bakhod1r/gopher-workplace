# Argmax over a count map

## Intuition

Count occurrences, then select the key with the highest count, breaking ties
deterministically (here, the smaller value):

```go
best, bestN, have := 0, 0, false
for v, n := range counts {
	if !have || n > bestN || (n == bestN && v < best) { best, bestN, have = v, n, true }
}
```

## Approach

1. Empty -> (0,false).
2. Count occurrences into a map.
3. Track best value: higher count wins; equal count keeps the smaller value.
4. Return (best,true).

## Solution

```go
func Mode(xs []int) (int, bool) {
	if len(xs) == 0 {
		return 0, false
	}
	counts := map[int]int{}
	for _, v := range xs {
		counts[v]++
	}
	best, bestCount := 0, -1
	for v, c := range counts {
		if c > bestCount || (c == bestCount && v < best) {
			best, bestCount = v, c
		}
	}
	return best, true
}
```

## Walkthrough

[1,1,2,2] counts {1:2,2:2}. Both count 2; smaller value 1 wins -> (1,true).

## Pitfalls

- Map iteration order is random — never rely on it; encode the tie-break.
- Return `ok=false` for empty rather than a fake value.
- Initialize the "best" carefully (use a `have` flag).

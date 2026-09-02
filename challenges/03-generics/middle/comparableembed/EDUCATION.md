# Embedding comparable

## Intuition

Writing both makes the intent explicit: keys must support `==` *and* be one of the shapes the storage layer understands.

## Approach

1. Allocate the map.
2. Increment the tally per element.
3. Return the map.

## Solution

```go
func Tally[K Key](s []K) map[K]int {
	out := make(map[K]int, len(s))
	for _, v := range s {
		out[v]++
	}
	return out
}
```

## Walkthrough

`Tally([]int{1,1,2})` counts `1` twice and `2` once; a struct key would not compile at all.

## Pitfalls

- Using bare `comparable` and letting struct keys through.
- Dropping the embed and assuming the set implies it — true here, but the intent is lost.
- Returning distinct-key counts instead of per-key tallies.

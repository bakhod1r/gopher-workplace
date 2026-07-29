# Map as an index

## Intuition

Walk once; for each element look up whether its complement was already seen. A
map from value to index gives O(1) lookups:

```go
seen := make(map[int]int)
for i, x := range xs {
	if j, ok := seen[target-x]; ok { return j, i, true }
	seen[x] = i
}
return 0, 0, false
```

## Approach

1. Map seen value -> earliest index.
2. For each element, look up complement target-v.
3. If found, return (storedIndex, currentIndex, true).
4. Otherwise record v's index if unseen.
5. Return false if no pair.

## Solution

```go
func TwoSum(xs []int, target int) (i, j int, ok bool) {
	seen := map[int]int{}
	for idx, v := range xs {
		if p, found := seen[target-v]; found {
			return p, idx, true
		}
		if _, dup := seen[v]; !dup {
			seen[v] = idx
		}
	}
	return 0, 0, false
}
```

## Walkthrough

[3,2,4] t=6: 3 store; 2 store; 4 complement 2 found at index1 -> (1,2,true).

## Pitfalls

- Check for the complement **before** inserting the current element, or you may
  pair an element with itself.
- The map stores the earliest index of each value.
- Returns the first pair found; multiple pairs may exist.

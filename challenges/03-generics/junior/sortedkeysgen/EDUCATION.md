# Sorted Keys

## Intuition

`cmp.Ordered` is a strictly stronger promise than `comparable`, so constraining `K` this way keeps the map type legal while adding the `<` needed to sort.

## Approach

1. Collect every key into `out`.
2. Sort `out` with `<`.
3. Return `out`.

## Solution

```go
func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	out := make([]K, 0, len(m))
	for k := range m {
		out = append(out, k)
	}
	for i := 1; i < len(out); i++ {
		for j := i; j > 0 && out[j] < out[j-1]; j-- {
			out[j], out[j-1] = out[j-1], out[j]
		}
	}
	return out
}
```

## Walkthrough

`SortedKeys(map[int]bool{3: true, 1: false})` may collect `[3 1]` or `[1 3]`; sorting makes both runs return `[1 3]`.

## Pitfalls

- Declaring `K comparable`, which cannot be sorted with `<`.
- Assuming map range already yields sorted keys.
- Sorting the values instead of the keys.

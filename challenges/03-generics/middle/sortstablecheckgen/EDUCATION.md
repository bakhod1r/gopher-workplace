# Verifying Stability

## Intuition

Testing stability requires knowing the input order, which is why the helper needs the original slice: the sorted output alone cannot reveal it.

## Approach

1. Reject any descending key step.
2. Record each element's first position in the original.
3. Reject equal-key neighbours whose original positions are inverted.

## Solution

```go
func IsStableBy[T comparable, K cmp.Ordered](s, original []T, key func(T) K) bool {
	for i := 1; i < len(s); i++ {
		if key(s[i]) < key(s[i-1]) {
			return false
		}
	}
	pos := make(map[T]int, len(original))
	for i, v := range original {
		if _, ok := pos[v]; !ok {
			pos[v] = i
		}
	}
	for i := 1; i < len(s); i++ {
		if key(s[i]) == key(s[i-1]) && pos[s[i]] < pos[s[i-1]] {
			return false
		}
	}
	return true
}
```

## Walkthrough

Two rows with the same key that swapped places relative to the original make the second check fail.

## Pitfalls

- Checking only that the result is sorted.
- Comparing positions across different keys, which says nothing about stability.
- Assuming the sorted slice alone carries the input order.

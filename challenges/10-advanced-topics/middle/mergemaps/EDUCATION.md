# Merge Into The Map You Were Given

## Intuition

The map parameter is a handle to the caller's table, so merging in place is both cheaper and visible. The only care needed is counting: the write itself cannot tell you whether the key was there.

## Approach

1. Return 0 for a nil `dst`.
2. For each entry, test presence with comma-ok, then write.
3. Return the count of keys that were absent.

## Solution

```go
// Merge copies every entry of src into dst, overwriting existing keys,
// and returns how many keys were newly added.
//
// dst is modified in place; a nil dst adds nothing.
//
// Examples:
//
// 	Merge(dst, map[string]int{"a": 1}) => 1 when dst lacked "a"
func Merge(dst, src map[string]int) int {
	if dst == nil {
		return 0
	}
	added := 0
	for k, v := range src {
		if _, ok := dst[k]; !ok {
			added++
		}
		dst[k] = v
	}
	return added
}
```

## Walkthrough

Merging {"a":9,"b":2} into {"a":1}: "a" is present so the count stays 0 and the value is overwritten; "b" is absent, so the count becomes 1.

## Pitfalls

- Counting `len(dst)` before and after — correct, and it hides the intent.
- Checking `dst[k] != 0`, which mistakes a stored zero for an absent key.

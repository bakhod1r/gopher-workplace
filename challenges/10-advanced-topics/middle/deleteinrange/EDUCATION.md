# Delete While You Range

## Intuition

Go's map iteration is explicitly specified to tolerate deletion: entries removed before they are reached are not produced. The defensive two-pass version is a habit carried over from languages where it would crash.

## Approach

1. Range the keys.
2. Delete the even ones and count them.

## Solution

```go
// RemoveEven deletes every entry whose key is even and returns how many
// were removed.
//
// Deleting during a range is defined: an entry not yet reached that is
// deleted will not be produced.
//
// Examples:
//
// 	RemoveEven(map[int]int{1: 1, 2: 2}) => 1
func RemoveEven(m map[int]int) int {
	removed := 0
	for k := range m {
		if k%2 == 0 {
			delete(m, k)
			removed++
		}
	}
	return removed
}
```

## Walkthrough

Over a 1000-entry map, 500 keys are deleted as the loop reaches them, and the loop still visits each surviving key exactly once.

## Pitfalls

- Adding entries during iteration, which *is* unspecified — new keys may or may not be produced.
- `k%2 == 0` is correct for negatives too; `k&1 == 0` is as well, but `k%2 == 1` is not a test for odd.

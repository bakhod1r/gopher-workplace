# Bool sets store true

## Intuition

When a `map[T]bool` is used as a set and membership is tested by the value
(`if inB[x]`), members must be stored as **true**:

```go
inB[x] = true
```

Storing `false` makes every lookup report "absent".

## Approach

1. Bug: inB[x] = false stores false for every b element; then inB[x] reads false and membership always fails. 2. Presence in a map and the stored value are different; here the code relies on the value being true. 3. Fix: inB[x] = true.

## Solution

```go
func Intersect(a, b []int) []int {
	inB := make(map[int]bool)
	for _, x := range b {
		inB[x] = true
	}
	out := []int{}
	seen := make(map[int]bool)
	for _, x := range a {
		if inB[x] && !seen[x] {
			out = append(out, x)
			seen[x] = true
		}
	}
	return out
}
```

## Walkthrough

b=[2,3,4] populates inB with 2,3,4 -> true. For a's 2 and 3, inB[x]=true and unseen -> appended. With false stored, the &&-guard is always false, output empty.

## Pitfalls

- With value-based tests, the stored bool must be `true`.
- `map[T]struct{}` uses no value memory and tests via `ok`.
- Reading a missing key gives `false` — indistinguishable from a stored `false`.

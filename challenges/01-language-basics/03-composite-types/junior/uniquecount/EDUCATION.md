# Sets from maps

## Intuition

Go has no set type; you use a map whose keys are the members. The value carries
no information, so `struct{}` (zero bytes) is idiomatic:

```go
seen := make(map[int]struct{})
for _, x := range xs { seen[x] = struct{}{} }
return len(seen)
```

## Approach

1. make a set map[int]struct{} (or map[int]bool).
2. Range xs, inserting each value; duplicates collapse to one key.
3. Return len(seen), the count of distinct values.

## Solution

```go
func Distinct(xs []int) int {
	seen := make(map[int]struct{})
	for _, x := range xs {
		seen[x] = struct{}{}
	}
	return len(seen)
}
```

## Walkthrough

Distinct([1,2,2,3,3,3]): set becomes {1,2,3}; len is 3.

## Pitfalls

- `map[int]struct{}` uses no memory for values; `map[int]bool` also works but
  stores a byte.
- Membership test is `_, ok := seen[x]`.
- The element type must be comparable to be a map key.

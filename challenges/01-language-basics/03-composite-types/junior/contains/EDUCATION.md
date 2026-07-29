# Membership by linear scan

## Intuition

For an unsorted slice, membership is a linear scan with an early return:

```go
for _, x := range xs {
	if x == target { return true }
}
return false
```

## Approach

1. Range over xs.
2. Compare each element to target; return true on the first match.
3. If the loop finishes with no match, return false.

## Solution

```go
func Contains(xs []string, target string) bool {
	for _, x := range xs {
		if x == target {
			return true
		}
	}
	return false
}
```

## Walkthrough

Contains(["a","b","c"],"c"): "a"!="c", "b"!="c", "c"=="c" -> return true.

## Pitfalls

- Return `false` only after the whole loop — not inside it.
- `==` must be defined for the element type (strings, numbers, comparable
  structs).
- Ranging nil is safe and yields `false`.

# Summarise The Leaves

## Intuition

Only leaves are distinct events; wrappers and join nodes are structure. Counting them too would inflate every total by the depth of the tree.

## Approach

1. Create the map once and pass it down the recursion.
2. Recurse through both unwrap shapes without counting.
3. Increment only at a leaf.

## Solution

```go
// Summary:
out := make(map[string]int)
count(err, out)
return out

// count:
if err == nil {
	return
}
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	for _, e := range joined.Unwrap() {
		count(e, out)
	}
	return
}
if wrapped, ok := err.(interface{ Unwrap() error }); ok {
	count(wrapped.Unwrap(), out)
	return
}
out[err.Error()]++
```

## Walkthrough

`fmt.Errorf("x: %w", ErrA)` contributes only `"a"`, because the wrapper is traversed and never counted.

## Pitfalls

- Counting every node, inflating totals with wrapper messages.
- Returning a nil map for a nil error.
- Building a fresh map at each recursion level and losing counts.

# Bounded Matching

## Intuition

`errors.Is` walks until the chain ends, which is fine for chains you build and dangerous for chains you receive. A depth cap makes the cost of a match bounded by policy rather than by input.

## Approach

1. Return false for a non-positive `max` or a nil argument.
2. Compare each link with `==` or a single-level `errors.Is`.
3. Stop after `max` links.

## Solution

```go
if max <= 0 || err == nil || target == nil {
	return false
}
for i := 0; i < max && err != nil; i++ {
	if err == target {
		return true
	}
	err = errors.Unwrap(err)
}
return false
```

## Walkthrough

The four-link chain matches at depth 4, so a cap of 3 reports false even though the sentinel is present deeper down.

## Pitfalls

- Calling `errors.Is` on the whole chain, which ignores the cap entirely.
- Counting from zero and allowing one extra link.
- Treating a zero cap as unlimited.

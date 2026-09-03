# Match Any Sentinel

## Intuition

Classification usually means "is this one of these?" A variadic wrapper around `errors.Is` keeps every call site free of repeated `||` chains.

## Approach

1. Range over the targets.
2. Return true on the first `errors.Is` match.
3. Return false after the loop.

## Solution

```go
for _, target := range targets {
	if errors.Is(err, target) {
		return true
	}
}
return false
```

## Walkthrough

With no targets the loop body never runs and the answer is false — nothing was asked about.

## Pitfalls

- Comparing `err == target`, which fails on wrapped errors.
- Returning false inside the loop, so only the first target is checked.
- Treating an empty target list as a match.

# Classify By Table

## Intuition

Turning a policy into data makes it reviewable and testable on its own. Ordering matters when rules overlap, which is why a slice beats a map here.

## Approach

1. Return 200 for nil.
2. Return the first rule whose sentinel matches.
3. Return 500 as the default.

## Solution

```go
if err == nil {
	return 200
}
for _, r := range table {
	if errors.Is(err, r.Err) {
		return r.Code
	}
}
return 500
```

## Walkthrough

An empty table matches nothing, so even a known sentinel falls through to 500 — the policy is entirely in the data.

## Pitfalls

- Using a `map[error]int`, which loses order and fails on wrapped errors.
- Returning 0 for unmatched errors.
- Checking nil after the loop, so `errors.Is(nil, …)` runs first.

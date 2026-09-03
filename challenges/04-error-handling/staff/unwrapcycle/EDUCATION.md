# Survive A Cycle

## Intuition

Nothing in the `error` interface promises that unwrapping terminates. Any traversal that runs over values a caller can construct needs its own termination guarantee.

## Approach

1. Track seen errors in a `map[error]bool`.
2. Append each message and stop when an error repeats.
3. Unwrap one level per iteration.

## Solution

```go
var out []string
seen := make(map[error]bool)
for err != nil {
	if seen[err] {
		break
	}
	seen[err] = true
	out = append(out, err.Error())
	err = errors.Unwrap(err)
}
return out
```

## Walkthrough

The self-wrapping error is recorded on the first pass, so the second iteration sees it in the map and stops with a single entry.

## Pitfalls

- Trusting the chain to end, which hangs on a cycle.
- Deduplicating by message, which truncates legitimate repeated text.
- Using a map key type that panics for uncomparable errors.

# Flattening Recursion For The Picture

## Intuition

Walk the stack keeping a frame only when it is not a repeat of the one you just kept. That is run-length encoding with the counts thrown away.

## Approach

1. Allocate a result with a capacity hint.
2. Append a frame when the result is empty or its last element differs.
3. `Depth` subtracts the collapsed length from the original.

## Solution

```go
func Collapse(stack []string) []string {
	out := make([]string, 0, len(stack))
	for _, fn := range stack {
		if len(out) == 0 || out[len(out)-1] != fn {
			out = append(out, fn)
		}
	}
	return out
}

func Depth(stack []string) int {
	return len(stack) - len(Collapse(stack))
}
```

## Walkthrough

`[a a b b a a]` becomes `[a b a]`: each run collapses, but the second `a` survives because a `b` separates it from the first.

## Pitfalls

- Deduplicating with a set, which turns mutual recursion into a straight line.
- Collapsing in place over the input slice, corrupting the caller's stack.
- Collapsing before summing values, which merges distinct paths that should have stayed apart.

# Only The Frame On The CPU

## Intuition

A stack sample is a snapshot of who called whom. Exactly one of those frames held the CPU: the last one.

## Approach

1. Skip junk samples.
2. Take the last frame and add the value to its total.

## Solution

```go
func SelfTime(samples []Sample) map[string]int64 {
	out := make(map[string]int64)
	for _, s := range samples {
		if s.Value <= 0 || len(s.Stack) == 0 {
			continue
		}
		out[s.Stack[len(s.Stack)-1]] += s.Value
	}
	return out
}
```

## Walkthrough

In the second example `a` is a leaf in two different stacks, so its self time is `3+4=7`, while `b` is a leaf only once and gets `1`.

## Pitfalls

- Indexing `[0]`, which credits the caller and inverts the whole profile.
- Indexing before the length guard, panicking on an empty stack.
- Summing self times and expecting them to exceed the profile total — that only happens with cum.

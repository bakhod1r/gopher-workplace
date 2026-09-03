# The Recursive Function That Owns 4000% Of The CPU

## Intuition

The set exists to answer one question — "have I already credited this frame for this sample?" — and the answer has to change what happens next.

## Approach

1. Skip the frame when it is already in the set.

## Solution

```go
for _, fn := range s.Stack {
	if seen[fn] {
		continue
	}
	seen[fn] = true
	out[fn] += s.Value
}
```

## Walkthrough

With the bug, a 200-frame recursion credits `descend` two hundred times for one sample. The seen-set is populated correctly and consulted never, so it looks like the fix is present in every code review.

## Pitfalls

- Hoisting the seen-set out of the sample loop, which credits each function once for the entire profile.
- Deduplicating the stack before the loop, which also destroys the caller-callee order other views need.
- Accepting a cum figure that exceeds the wall clock by orders of magnitude.

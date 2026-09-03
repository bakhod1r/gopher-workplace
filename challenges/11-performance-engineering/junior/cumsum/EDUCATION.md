# The Cumulative Column

## Intuition

A sample means "these frames were all on the stack at this instant". Each of them was responsible for that instant — once, no matter how many frames deep the recursion went.

## Approach

1. For each valid sample, allocate a seen-set.
2. Walk the stack, skipping frames already seen for this sample.
3. Add the value to the running total for each newly seen frame.

## Solution

```go
func CumSum(samples []Sample) map[string]int64 {
	out := make(map[string]int64)
	for _, s := range samples {
		if s.Value <= 0 || len(s.Stack) == 0 {
			continue
		}
		seen := make(map[string]bool, len(s.Stack))
		for _, fn := range s.Stack {
			if seen[fn] {
				continue
			}
			seen[fn] = true
			out[fn] += s.Value
		}
	}
	return out
}
```

## Walkthrough

For `[main rec rec rec]` with value 6, `rec` is credited on the first frame and skipped on the next two, giving `6` rather than `18`.

## Pitfalls

- Hoisting the seen-set out of the sample loop, which credits each function at most once for the whole profile.
- Charging only the leaf, which is flat time, not cum.
- Forgetting that cum totals legitimately sum to more than the profile's wall time.

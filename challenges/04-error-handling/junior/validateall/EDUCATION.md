# Collect All Failures

## Intuition

Aborting on the first failure is right for a computation, wrong for a validation report. Gathering every failure means the user fixes all the rows in one pass.

## Approach

1. Declare `var out []error`.
2. Append `ErrNegative` for every negative entry.
3. Return `out` — nil when nothing was appended.

## Solution

```go
var out []error
for _, n := range nums {
	if n < 0 {
		out = append(out, ErrNegative)
	}
}
return out
```

## Walkthrough

For `[]int{1, -2, -3}` the loop appends twice; for `[]int{1, 2}` nothing is appended and `out` is still nil.

## Pitfalls

- Initialising with `out := []error{}` — the no-failure case must be nil.
- Returning early on the first failure, defeating the purpose.
- Pre-allocating `make([]error, len(nums))`, which leaves nil holes.

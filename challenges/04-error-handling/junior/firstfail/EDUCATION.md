# Index Of First Failure

## Intuition

A search returns two things: where the match is, and whether there was one at all. Pairing an index with an error means the caller cannot use `-1` as a real position by accident.

## Approach

1. Range with both index and value.
2. Return `i, nil` at the first non-nil entry.
3. Return `-1, ErrNoFailure` after the loop.

## Solution

```go
for i, err := range errs {
	if err != nil {
		return i, nil
	}
}
return -1, ErrNoFailure
```

## Walkthrough

For `[]error{nil, ErrStep}` the first iteration skips, and the second returns `1, nil`.

## Pitfalls

- Returning `0` instead of `-1` when nothing failed — index 0 is a real position.
- Returning `-1, nil`, so callers must check the number instead of the error.
- Using a manual counter that drifts out of step with the loop.

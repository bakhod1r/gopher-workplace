# Most Recent Failure

## Intuition

Finding the *last* match cannot stop early — every remaining element might replace the answer. The result is carried in a variable declared outside the loop.

## Approach

1. Declare `var last error` before the loop.
2. Assign every non-nil entry to it.
3. Return `last` after the loop.

## Solution

```go
var last error
for _, err := range errs {
	if err != nil {
		last = err
	}
}
return last
```

## Walkthrough

For `[]error{ErrA, nil, ErrB}`: `last` becomes `ErrA`, is skipped for nil, then becomes `ErrB`.

## Pitfalls

- Returning inside the loop, which yields the first failure.
- Assigning unconditionally, so a trailing nil erases a real failure.
- Declaring `last` inside the loop, losing it each pass.

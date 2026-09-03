# Count By Class

## Intuition

Aggregating by class turns a flood of individual failures into a signal. The classification uses `errors.Is`, so annotated failures still land in the right bucket.

## Approach

1. Create the map up front.
2. Skip nil entries.
3. Increment the bucket chosen by `errors.Is`, defaulting to `"other"`.

## Solution

```go
out := make(map[string]int)
for _, err := range errs {
	if err == nil {
		continue
	}
	switch {
	case errors.Is(err, ErrTimeout):
		out["timeout"]++
	case errors.Is(err, ErrDenied):
		out["denied"]++
	default:
		out["other"]++
	}
}
return out
```

## Walkthrough

The wrapped `ErrDenied` is counted under `"denied"` because `errors.Is` sees through the annotation.

## Pitfalls

- Returning a nil map, which is not equal to an empty one.
- Counting nil entries as `"other"`.
- Comparing with `==`, so wrapped failures all fall into `"other"`.

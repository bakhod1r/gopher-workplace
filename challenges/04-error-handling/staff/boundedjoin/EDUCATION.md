# Join With A Cap

## Intuition

An unbounded aggregate error is a memory and log-volume incident of its own. Keeping a sample plus an accurate remainder count preserves the information that actually drives a decision.

## Approach

1. Collect non-nil errors up to `max`, counting the rest.
2. Append a summary error when the count is positive.
3. Return `errors.Join` of the result.

## Solution

```go
var kept []error
dropped := 0
for _, err := range errs {
	if err == nil {
		continue
	}
	if len(kept) < max {
		kept = append(kept, err)
		continue
	}
	dropped++
}
if dropped > 0 {
	kept = append(kept, fmt.Errorf("and %d more", dropped))
}
return errors.Join(kept...)
```

## Walkthrough

With a cap of 1 and two real failures, one is kept and the summary reports exactly one dropped — the nils never count.

## Pitfalls

- Counting nil entries as dropped.
- Adding the summary when nothing was dropped.
- Joining everything first and truncating the message afterwards.

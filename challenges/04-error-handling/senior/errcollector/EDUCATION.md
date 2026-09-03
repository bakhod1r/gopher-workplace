# Bounded Error Collector

## Intuition

Collecting every failure sounds thorough until a bad deploy produces a million of them. Keeping a sample plus an exact count preserves both signal and memory.

## Approach

1. Ignore nil in `Add`.
2. Increment the count, then store only while under the limit.
3. Join the stored slice in `Err`.

## Solution

```go
// Add:
if err == nil {
	return
}
c.count++
if len(c.stored) < c.Limit {
	c.stored = append(c.stored, err)
}

// Count:
return c.count

// Err:
return errors.Join(c.stored...)
```

## Walkthrough

With a limit of 2 and three failures, the count is 3 while only the first two remain matchable.

## Pitfalls

- Counting nil adds.
- Storing first and trimming later, defeating the memory bound.
- Returning a non-nil join when nothing was stored.

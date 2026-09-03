# Stable Ordering

## Intuition

Nondeterministic ordering makes two identical outcomes look like a change. Sorting by message gives a total order that survives concurrency, map iteration and scheduling.

## Approach

1. Collect the non-nil errors into a new slice.
2. Sort by `Error()`.
3. Join the sorted slice.

## Solution

```go
var kept []error
for _, err := range errs {
	if err != nil {
		kept = append(kept, err)
	}
}
sort.Slice(kept, func(i, j int) bool {
	return kept[i].Error() < kept[j].Error()
})
return errors.Join(kept...)
```

## Walkthrough

Duplicates are preserved, so two occurrences of `"a"` produce two lines — deduplication is a separate decision.

## Pitfalls

- Sorting the variadic slice in place, mutating the caller's array when it was expanded from a named slice.
- Deduplicating while sorting, silently changing the counts.
- Sorting after joining, which is no longer possible.

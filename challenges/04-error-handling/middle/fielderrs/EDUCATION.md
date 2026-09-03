# Errors By Field

## Intuition

Go randomises map iteration on purpose, so any output built straight from a map range is unstable. Sorting the keys is what makes the error message testable.

## Approach

1. Collect the keys into a slice and sort them.
2. Wrap each non-nil entry with its field name.
3. Return `errors.Join` of the wrappers.

## Solution

```go
keys := make([]string, 0, len(m))
for k := range m {
	keys = append(keys, k)
}
sort.Strings(keys)

var errs []error
for _, k := range keys {
	if err := m[k]; err != nil {
		errs = append(errs, fmt.Errorf("%s: %w", k, err))
	}
}
return errors.Join(errs...)
```

## Walkthrough

`errors.Join` separates its members with `\n`, so two fields render as two lines in sorted key order.

## Pitfalls

- Ranging the map directly, producing a message that changes between runs.
- Joining with `"; "` by hand and losing `errors.Is` support.
- Including nil entries as empty lines.

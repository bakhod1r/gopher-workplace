# Parse What You Can

## Intuition

Partial success is a real outcome. Returning both the parsed rows and the joined failures lets the caller decide whether to proceed, without forcing an all-or-nothing contract.

## Approach

1. Range with the index.
2. Append successes; wrap failures with the line number.
3. Return the values and `errors.Join` of the failures.

## Solution

```go
var out []int
var errs []error
for i, line := range lines {
	n, err := strconv.Atoi(line)
	if err != nil {
		errs = append(errs, fmt.Errorf("line %d: %w", i, err))
		continue
	}
	out = append(out, n)
}
return out, errors.Join(errs...)
```

## Walkthrough

For `["1", "x", "3"]` the values `1` and `3` are returned alongside a single joined failure naming line 1.

## Pitfalls

- Returning nil values when anything failed, discarding usable rows.
- Aborting at the first bad line.
- Numbering by the output slice, so the reported line drifts after each failure.

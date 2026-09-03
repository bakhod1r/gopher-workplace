# Annotate Every Failure

## Intuition

Filtering changes positions. If the annotation is taken from the output slice's length, every skipped success shifts the numbering and the labels stop pointing at real jobs.

## Approach

1. Declare `var out []error`.
2. Range with the index, skipping nil entries.
3. Append `fmt.Errorf("job %d: %w", i, err)`.

## Solution

```go
var out []error
for i, err := range errs {
	if err == nil {
		continue
	}
	out = append(out, fmt.Errorf("job %d: %w", i, err))
}
return out
```

## Walkthrough

For `[nil, ErrJob, nil, ErrJob]` the outputs are labelled 1 and 3 — the indexes in the input, not 0 and 1.

## Pitfalls

- Using `len(out)` as the job number, renumbering after every skip.
- Appending nil entries and wrapping them.
- Returning an empty non-nil slice when nothing failed.

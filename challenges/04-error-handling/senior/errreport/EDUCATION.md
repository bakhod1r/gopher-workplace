# Render A Report

## Intuition

Rendering and matching are different jobs. The report walks the joined members for structure rather than splitting the combined message, so a member containing a newline cannot corrupt the output.

## Approach

1. Return `""` for nil.
2. Collect the members via `Unwrap() []error`, or the single error.
3. Format each with a 1-based number and join with newlines.

## Solution

```go
if err == nil {
	return ""
}
members := []error{err}
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	members = joined.Unwrap()
}
lines := make([]string, 0, len(members))
for i, e := range members {
	lines = append(lines, fmt.Sprintf("%d. %s", i+1, e.Error()))
}
return strings.Join(lines, "\n")
```

## Walkthrough

`errors.Join(nil, ErrB)` holds one member, so the report is a single line numbered 1.

## Pitfalls

- Splitting `err.Error()` on newlines, which breaks on multi-line members.
- Numbering from 0.
- Returning `"0 errors"` or similar for nil instead of an empty string.

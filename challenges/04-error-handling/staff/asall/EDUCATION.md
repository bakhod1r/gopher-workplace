# Collect Every Typed Error

## Intuition

`errors.As` answers "is there one?" and hands back the first. Reporting every offending field is a different question, and the standard library deliberately does not answer it.

## Approach

1. Return nil for nil.
2. Append the node when it asserts to `*FieldError`.
3. Recurse into joined branches, then into a wrapped child.

## Solution

```go
if err == nil {
	return nil
}
var out []*FieldError
if fe, ok := err.(*FieldError); ok {
	out = append(out, fe)
}
if joined, ok := err.(interface{ Unwrap() []error }); ok {
	for _, e := range joined.Unwrap() {
		out = append(out, All(e)...)
	}
	return out
}
if wrapped, ok := err.(interface{ Unwrap() error }); ok {
	out = append(out, All(wrapped.Unwrap())...)
}
return out
```

## Walkthrough

The nested case finds `a` under a `fmt.Errorf` wrapper and `b` inside an inner join, returning both in tree order.

## Pitfalls

- Calling `errors.As` in a loop, which keeps finding the same first match.
- Returning after the first hit.
- Forgetting that a matching node may itself have children.

# List Joined Messages

## Intuition

`errors.Join` renders its members newline-separated, but that is a formatting detail. Asking the error for its members through `Unwrap() []error` is what the interface actually promises.

## Approach

1. Return nil for a nil error.
2. Assert to `interface{ Unwrap() []error }`.
3. Collect each member's message, or return the single message.

## Solution

```go
if err == nil {
	return nil
}
joined, ok := err.(interface{ Unwrap() []error })
if !ok {
	return []string{err.Error()}
}
out := make([]string, 0, len(joined.Unwrap()))
for _, e := range joined.Unwrap() {
	out = append(out, e.Error())
}
return out
```

## Walkthrough

`errors.Join(nil, ErrB)` holds one member, so the result is a single-element slice rather than an empty first entry.

## Pitfalls

- Splitting `err.Error()` on `"\n"`, which breaks when a member's own message contains a newline.
- Returning nil for a non-joined error instead of a one-element slice.
- Ignoring that `Join` already dropped nil members.

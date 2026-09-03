# Annotate And Wrap

## Intuition

Context makes an error diagnosable; wrapping makes it still matchable. `%w` gives you both — the message gains a prefix and the original stays reachable through the chain.

## Approach

1. Return nil immediately for a nil error.
2. Build the message with `fmt.Errorf` and a `%w` verb.
3. Return the wrapper.

## Solution

```go
if err == nil {
	return nil
}
return fmt.Errorf("%s: %w", op, err)
```

## Walkthrough

`Wrap("handler", Wrap("read", ErrDisk))` builds a two-link chain; unwrapping twice reaches `ErrDisk`, so `errors.Is` succeeds at any depth.

## Pitfalls

- Using `%v`, which flattens the error to text and breaks `errors.Is`.
- Wrapping nil, which turns success into a non-nil error.
- Using `errors.New(op + ": " + err.Error())` — same message, no chain.

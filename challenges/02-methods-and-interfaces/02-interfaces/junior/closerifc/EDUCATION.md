# Closer Interface

## Intuition

`Close() error` is the smallest useful contract in Go. Anything holding a resource can satisfy it, and generic shutdown code needs nothing else.

## Approach

1. In `Close`, return `ErrAlreadyClosed` when `f.Closed` is already true.
2. Otherwise set `f.Closed = true` and return `nil`.
3. In `CloseAll`, range over the slice and return the first non-nil error; return `nil` at the end.

## Solution

```go
func (f *File) Close() error {
	if f.Closed {
		return ErrAlreadyClosed
	}
	f.Closed = true
	return nil
}

func CloseAll(cs []Closer) error {
	for _, c := range cs {
		if err := c.Close(); err != nil {
			return err
		}
	}
	return nil
}
```

## Walkthrough

With `[]Closer{bad, good}` where `bad` is already closed: the first call returns `ErrAlreadyClosed`, `CloseAll` returns immediately, and `good` is never touched — which the test checks.

## Pitfalls

- Value receiver on `Close` — the mutation is lost and the double close is never detected.
- Collecting errors but returning the *last* one instead of the first.
- Continuing the loop after an error when the spec says stop.

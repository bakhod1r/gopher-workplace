# Boolean Mutation

## Intuition

Toggling is the simplest mutation: flip a boolean. The key learning is that
this requires a pointer receiver.

## Approach

1. Negate `s.On`.

## Solution

```go
func (s *Switch) Toggle() {
	s.On = !s.On
}
```

## Walkthrough

For `Switch{On: false}`:
- `!false` = `true`.
- `s.On` is now `true`.

## Pitfalls

- Value receiver → the flip is lost.
- Using `if/else` instead of `!` — correct but verbose.

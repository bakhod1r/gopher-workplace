# Why Push Needs a Pointer Receiver

## Intuition

`append` sometimes returns a *new* backing array. If `Push` used a value
receiver, the new slice header (with updated length and possibly a new pointer)
would be stored in the copy, not the caller's struct. A pointer receiver ensures
the caller's `Items` field gets updated.

## Approach

1. `append(s.Items, v)` and assign back to `s.Items`.

## Solution

```go
func (s *Stack) Push(v int) {
	s.Items = append(s.Items, v)
}
```

## Walkthrough

For `s := Stack{}; s.Push(5)`:
- `append(nil, 5)` → `[]int{5}`.
- `s.Items` now points to the new slice.

## Pitfalls

- Value receiver `(s Stack)` — `append` updates the copy's `Items`; caller's
  `Items` stays `nil`.
- This is the most common Go method bug: methods that `append` to a slice MUST
  use pointer receivers.

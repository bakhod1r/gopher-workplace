# Reset Methods

## Intuition

A `Reset` method is a common pattern: restore an object to its initial state.
Since it mutates the receiver, it needs a pointer receiver.

## Approach

1. Set `c.N` to 0.

## Solution

```go
func (c *Counter) Reset() {
	c.N = 0
}
```

## Walkthrough

For `Counter{42}`:
- `c.N = 0` → `c.N` is now 0.

## Pitfalls

- Value receiver `(c Counter)` would reset a copy — caller's counter unchanged.
- `*c = Counter{}` also works (resets entire struct) but is overkill for a
  single field.

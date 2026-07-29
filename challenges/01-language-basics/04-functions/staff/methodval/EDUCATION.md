# Method values snapshot the receiver

## Intuition

Binding a value-receiver method (`c.Get`) copies the receiver then; the bound function is independent of later changes to `c` — unless you re-bind.

## Approach

1. `c.Get` is a method **value**: it binds the receiver when created.
2. The stray reassignment `f = c.Get` is redundant/incorrect; remove it so the intended binding stands.

## Solution

```go
type counter struct{ n int }

func (c counter) Get() int { return c.n }

func BoundEarly(start int) func() int {
	c := counter{n: start}
	f := c.Get // method value: receiver copied now
	c.n = 999  // must NOT affect f (value receiver was copied)
	return f
}
```

## Walkthrough

The method value already captures the receiver's state at bind time; the extra line re-binds to a stale copy. Removing it returns the correctly bound closure.

## Pitfalls

- `f := c.Get` copies `c` now; mutating `c` afterwards doesn't change `f`.
- Re-binding after a mutation captures the newer receiver.

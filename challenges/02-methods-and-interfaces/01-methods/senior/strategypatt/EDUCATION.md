# Strategy Pattern

## Intuition

Strategy separates *what to do to each item* from *how to walk the items*. In
languages without first-class functions this needs an interface and a class per
algorithm; in Go the strategy is a `func` value, and the pattern shrinks to a
parameter.

## Approach

1. Iterate by index, so each slot is addressable.
2. Replace each element with the strategy's result.

## Solution

```go
func (c *Context) Process(strategy func(int) int) {
	for i := range c.Data {
		c.Data[i] = strategy(c.Data[i])
	}
}
```

## Walkthrough

`for i := range c.Data` yields indices 0, 1, 2. Each iteration reads
`c.Data[i]`, applies `double`, and stores the result back into the same slot.
Because a slice value holds a pointer to its backing array, the caller's
`c.Data` shows `[2 4 6]` — `reflect.DeepEqual` in the test confirms it.

## Pitfalls

- **`for _, v := range c.Data { v = strategy(v) }`.** `v` is a fresh copy each
  iteration; the assignment is discarded and the data is unchanged.
- **Building a new slice and not assigning it.** `out := make(...)` then
  forgetting `c.Data = out`.
- **Value receiver.** Writing through `c.Data[i]` would still work — the slice
  header is copied but the array is shared — yet `c.Data = out` would not. The
  pointer receiver keeps both styles correct.

## The one place a value receiver would still mutate

Slices, maps and channels are reference-like: a copy of the header points at the
same data. That is why in-place element writes escape a value receiver, while
reassigning the field does not. Relying on that distinction is a well-known
source of confusion — prefer the pointer receiver and be explicit.

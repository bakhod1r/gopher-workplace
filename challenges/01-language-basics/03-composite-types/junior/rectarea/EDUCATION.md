# Value-receiver methods

## Intuition

A method with a **value receiver** operates on a copy, so it can't mutate the
caller — ideal for pure computations that return new values:

```go
func (r Rect) Area() int      { return r.W * r.H }
func (r Rect) Scale(f int) Rect { return Rect{r.W*f, r.H*f} }
```

## Approach

1. Area: return r.W * r.H.
2. Scale: return a new Rect{W: r.W*factor, H: r.H*factor}.
3. Both use value receivers, so the caller's Rect is never mutated.

## Solution

```go
type Rect struct {
	W, H int
}

func (r Rect) Area() int {
	return r.W * r.H
}

func (r Rect) Scale(factor int) Rect {
	return Rect{W: r.W * factor, H: r.H * factor}
}
```

## Walkthrough

Rect{2,3}.Scale(2): builds Rect{2*2, 3*2} = Rect{4,6}; the receiver r is a copy, so the original stays {2,3}.

## Pitfalls

- A value receiver copies the struct; mutations inside are lost. Use a pointer
  receiver to mutate.
- Structs of comparable fields support `==` and can be map keys.
- Returning a new struct keeps the original untouched (no surprise mutation).

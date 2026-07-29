# Field access through struct pointers

## Intuition

Go implicitly dereferences a struct pointer for field selection, so `c.Count` reads/writes the pointee's field.

## Approach

1. `c` is a `*Cart`.
2. `c.Count++` — Go auto-dereferences the pointer for field access, so this mutates the caller's struct.

## Solution

```go
type Cart struct{ Count int }

func Grow(c *Cart) {
	c.Count++
}
```

## Walkthrough

`Grow(&c)` receives the address; `c.Count++` is shorthand for `(*c).Count++`, incrementing the caller's field.

## Pitfalls

- No need to write `(*c).Count`; `c.Count` is equivalent.
- A value receiver would edit a copy, not the caller's struct.

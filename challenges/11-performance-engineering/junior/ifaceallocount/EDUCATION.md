# Boxing Happens At The Boundary

## Intuition

An interface value is a type word and a data word. When the data is already a pointer, storing it is a move — no allocation, no copy of the struct.

## Approach

1. `Area` multiplies the two fields.
2. `TotalArea` ranges and accumulates.

## Solution

```go
func (r *Rect) Area() int {
	return r.W * r.H
}

func TotalArea(shapes []Shape) int {
	total := 0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}
```

## Walkthrough

The allocation, if any, happened when each `&Rect{...}` was placed in the slice. By the time `TotalArea` runs, every interface already holds a pointer, so the loop touches the allocator zero times.

## Pitfalls

- A value receiver plus `[]Shape{Rect{...}}`, which copies each struct to the heap on every conversion.
- Blaming the interface call for allocations that happened at the conversion site.
- Expecting `Rect` to satisfy `Shape` when the receiver is `*Rect`.

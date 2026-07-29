# Addresses of reused variables

## Intuition

A variable declared outside the loop has one address; taking `&v` each iteration stores the same pointer. A per-iteration declaration gives distinct addresses.

## Approach

1. A single hoisted `var v` is reused every iteration, so all pointers alias it.
2. Declare `v := xs[i]` **inside** the loop to get fresh storage each time.

## Solution

```go
func Pointers(xs []int) []*int {
	var out []*int
	for i := 0; i < len(xs); i++ {
		v := xs[i]
		out = append(out, &v)
	}
	return out
}
```

## Walkthrough

With the shared `v`, every `&v` points at the same cell holding the last value. A per-iteration `v` gives each pointer its own element value.

## Pitfalls

- `&v` of a hoisted `v` aliases one storage cell.
- Declare the variable inside the loop for distinct addresses.

# Aliasing vs copying slice views

## Intuition

Multiple slices of the same array share storage; copying into a new slice breaks that aliasing.

## Approach

1. `append([]int(nil), p[:]...)` copies into a new array — no aliasing.
2. Slice the array directly: `b := p[:]`.

## Solution

```go
func Views(p *[3]int) ([]int, []int) {
	a := p[:]
	b := p[:]
	return a, b
}
```

## Walkthrough

The bug's copy makes `b` independent, so `a[0] = 42` doesn't reach it. Slicing `p[:]` shares the backing array.

## Pitfalls

- `p[:]` aliases the array; `append([]int(nil), p[:]...)` copies it.
- Aliased views see each other's writes.

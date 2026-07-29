# Inserting into a slice

## Intuition

Insertion stitches three pieces: the head, the new value, and the tail:

```go
out := append([]int{}, xs[:i]...)
out = append(out, v)
out = append(out, xs[i:]...)
```

## Approach

1. Clamp i to [0, len(xs)].
2. Allocate result with cap len+1.
3. Append xs[:i], then v, then xs[i:].
4. Return result (input untouched).

## Solution

```go
func InsertAt(xs []int, i, v int) []int {
	if i < 0 {
		i = 0
	}
	if i > len(xs) {
		i = len(xs)
	}
	out := make([]int, 0, len(xs)+1)
	out = append(out, xs[:i]...)
	out = append(out, v)
	out = append(out, xs[i:]...)
	return out
}
```

## Walkthrough

[1,2,3] i=1 v=9: xs[:1]=[1], +9 -> [1,9], +xs[1:]=[2,3] -> [1,9,2,3].

## Pitfalls

- Building into a fresh slice avoids the aliasing clobber.
- `slices.Insert` (Go 1.21+) handles this correctly and efficiently.
- Clamp `i`; `xs[:i]` panics for `i > len`.

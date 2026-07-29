# In-place filtering aliases the input

## Intuition

`xs[:0]` is a length-0 slice over `xs`'s **same** backing array. Appending to it
writes over `xs[0], xs[1], ...`. That's the deliberate "filter in place" trick —
but only safe when you *want* to overwrite the input:

```go
out := []int{} // independent result; input untouched
```

## Approach

1. Bug: out := xs[:0] reuses xs's backing array; append(out, x) overwrites xs's leading elements as it filters, corrupting the input. 2. Fix: allocate a fresh slice out := []int{} (or make with cap). 3. Appending to an independent slice leaves xs intact.

## Solution

```go
func Evens(xs []int) []int {
	out := []int{}
	for _, x := range xs {
		if x%2 == 0 {
			out = append(out, x)
		}
	}
	return out
}
```

## Walkthrough

xs=[1,2,3,4], out=xs[:0]. append 2 writes xs[0]=2, append 4 writes xs[1]=4 -> xs corrupted to [2,4,3,4]. A fresh out collects [2,4] without touching xs.

## Pitfalls

- `xs[:0]` shares memory; a fresh slice does not.
- The in-place version (used by `slices.DeleteFunc`) mutates and returns the
  input.
- Reading `xs` after an in-place filter over it is undefined-ish — the front is
  overwritten.

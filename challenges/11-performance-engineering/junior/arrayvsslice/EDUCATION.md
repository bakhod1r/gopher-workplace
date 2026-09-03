# Arrays Copy, Slices Point

## Intuition

An array behaves like a struct with numbered fields: assignment and argument passing copy the whole thing. A slice behaves like a pointer with bounds attached.

## Approach

1. Sum with a `range` in both cases.
2. `ZeroBlock` returns `Block{}`; the parameter copy is discarded.

## Solution

```go
func SumBlock(b Block) int {
	total := 0
	for _, v := range b {
		total += v
	}
	return total
}

func ZeroBlock(b Block) Block {
	return Block{}
}

func SumSlice(s []int) int {
	total := 0
	for _, v := range s {
		total += v
	}
	return total
}
```

## Walkthrough

Neither function allocates: the array copy lives on the stack, and the slice header points at memory the caller already owns.

## Pitfalls

- Passing large arrays by value in a hot path — that is a full copy per call.
- Expecting a mutation of an array parameter to reach the caller; it never does.
- Forgetting that `b[:]` escapes the array to the heap if the slice outlives the frame.

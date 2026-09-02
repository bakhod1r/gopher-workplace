# When Not To Use Generics

## Intuition

The dividing line is whether the type appears once or across a collection: interfaces cost one indirection, while boxing a whole slice costs an allocation per element.

## Approach

1. `WriteAll`: call `w.Write` for each line and count.
2. `WriteEach`: format each value first, then write and count.

## Solution

```go
func WriteAll(w Writer, lines []string) int {
	n := 0
	for _, l := range lines {
		w.Write(l)
		n++
	}
	return n
}

func WriteEach[T any](w Writer, vs []T, format func(T) string) int {
	n := 0
	for _, v := range vs {
		w.Write(format(v))
		n++
	}
	return n
}
```

## Walkthrough

`WriteEach(w, []int{1,2}, itoa)` keeps the caller's `[]int` intact, whereas an interface-based version would need `[]any`.

## Pitfalls

- Writing `WriteAll[T Writer](w T, ...)`, which adds a type parameter for nothing.
- Taking `[]any` in `WriteEach` and boxing every element.
- Returning the concrete writer type and leaking the implementation.

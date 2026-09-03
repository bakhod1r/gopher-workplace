# The Header Is A Copy, The Array Is Not

## Intuition

A slice value is a small struct describing a window onto an array. Copying the struct copies the window, not the glass.

## Approach

1. `Fill` ranges by index and assigns.
2. `AppendLocal` appends to the parameter and discards the result.

## Solution

```go
func Fill(s []int, v int) {
	for i := range s {
		s[i] = v
	}
}

func AppendLocal(s []int, v int) {
	_ = append(s, v)
}
```

## Walkthrough

`append` writes `v` into the array at index 3 — that write is real and shared — but the new length lives only in the callee's copy of the header, so the caller still sees `len(s) == 3`.

## Pitfalls

- Expecting `append` inside a function to be visible to the caller; return the slice or pass `*[]int`.
- Handing a slice with spare capacity to two functions that both append into it.
- Assuming `range s` on a nil slice panics; it iterates zero times.

# Reinterpreting structs as arrays

## Intuition

A struct of homogeneous fields has the same memory layout as an array of that type, so an unsafe reinterpret gives an indexable view of all fields.

## Approach

1. Reinterpreting the struct as an array exposes both fields.
2. The bug returns only `arr[0]`; add both: `arr[0] + arr[1]`.

## Solution

```go
import "unsafe"

type Pair struct {
	A int32
	B int32
}

func Sum(p *Pair) int32 {
	arr := (*[2]int32)(unsafe.Pointer(p))
	return arr[0] + arr[1]
}
```

## Walkthrough

`Pair{3,4}` viewed as `[2]int` is `[3 4]`; summing both elements gives 7 rather than just 3.

## Pitfalls

- The reinterpreted array has one element per field.
- Reading only arr[0] ignores the rest of the struct.

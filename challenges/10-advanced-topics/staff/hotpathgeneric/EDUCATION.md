# Type Parameters Instead Of Interface Boxing

## Intuition

`any` erases the type, so every value needs a heap word and a type word. A type parameter keeps the type, so the compiler emits code for the real element width and never touches the allocator.

## Approach

1. Declare an int64 accumulator.
2. Range the values, converting each to int64 before adding.
3. Return the total.

## Solution

```go
// Total sums vals.
//
// An `[]any` version would box every element. A type parameter gives the
// compiler the concrete type, so nothing is boxed and nothing escapes.
//
// Examples:
//
// 	Total([]int{1, 2, 3}) => 6
func Total[T ~int | ~int32 | ~int64](vals []T) int64 {
	var sum int64
	for _, v := range vals {
		sum += int64(v)
	}
	return sum
}
```

## Walkthrough

Summing 1024 ints via `[]any` allocates about 1024 times, since values above 255 do not hit the runtime's small-integer cache. The generic version allocates zero and is several times faster.

## Pitfalls

- Accumulating in `T` — an `[]int32` of large values overflows before the conversion.
- Adding `~uint` to the constraint; the conversion to int64 would then be lossy.

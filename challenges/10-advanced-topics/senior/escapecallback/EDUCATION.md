# The Accumulator That Escaped Through A Callback

## Intuition

Escape analysis is per-function and conservative across call boundaries. Once a closure is passed to a function the compiler cannot see into, it has to assume the func value may be stored — so the closure and everything it captures go to the heap.

## Approach

1. Drop the callback and range over `s` directly.
2. Accumulate into a local, which now provably does not escape.

## Solution

```go
// Each calls f for every element of s.
//
// It is a package-level variable, so the compiler cannot see which function
// runs and must assume the callback it is given escapes.
var Each = func(s []int, f func(int)) {
	for _, v := range s {
		f(v)
	}
}

// Sum returns the total of s.
//
// The helper Each is not inlinable, so any closure handed to it escapes —
// and with it everything the closure captures.
//
// Examples:
//
// 	Sum([]int{1, 2, 3}) => 6
func Sum(s []int) int64 {
	var total int64
	for _, v := range s {
		total += int64(v)
	}
	return total
}
```

## Walkthrough

The callback version allocates the closure object and moves `total` to the heap — two allocations, plus an indirect call per element. The direct loop keeps `total` in a register.

## Pitfalls

- Assuming all closures allocate; one passed to a directly-called, inlinable function usually does not.
- Turning `Each` back into a plain function instead of removing the callback — the helper is part of the fixture.

# What A Closure Drags Onto The Heap

## Intuition

A closure over a mutable local turns that local into shared state between two frames — the one that made it and every call of the closure. It cannot stay in a frame that is about to disappear, so it is allocated.

## Approach

1. Copy `start` into a local.
2. Return a function that reads it, increments it, and returns the old value.

## Solution

```go
// Counter returns a function that yields start, start+1, start+2 and so
// on, one value per call.
//
// The captured variable outlives Counter's frame, so it must live on the
// heap — that is what a closure over a mutable local costs.
//
// Examples:
//
// 	c := Counter(1); c(), c() => 1, 2
func Counter(start int) func() int {
	n := start
	return func() int {
		v := n
		n++
		return v
	}
}
```

## Walkthrough

`Counter(1)` allocates the captured `n` and the closure object. Each `c()` reads and bumps the same heap word, which is why two counters made from the same start stay independent but two calls of one counter do not.

## Pitfalls

- Incrementing before returning, which yields `start+1` first.
- Capturing the parameter directly is fine — it is a local too.

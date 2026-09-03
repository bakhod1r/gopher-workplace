# Deferred Cleanup That Waits For The Whole Loop

## Intuition

`defer` binds to the function, not the block. In a loop it builds a stack of pending calls that all fire at the return — so the peak resource count is the whole loop, and the order is reversed.

## Approach

1. Drop the `defer`.
2. Call `release(v)` directly at the end of each iteration.

## Solution

```go
// Process doubles each item and calls release with the item as soon as
// that item is finished.
//
// release returns the item's resources. Holding every item until the
// function returns is what makes a batch job run out of them.
//
// Examples:
//
// 	Process([]int{1, 2}, rel) => []int{2, 4}, rel called after each item
func Process(items []int, release func(int)) []int {
	out := make([]int, 0, len(items))
	for _, v := range items {
		out = append(out, v*2)
		release(v)
	}
	return out
}
```

## Walkthrough

With `defer`, releases for 1,2,3 run after the return in the order 3,2,1, and all three items are held at once. Called inline, item 1 is released before item 2 is touched.

## Pitfalls

- Wrapping the body in a closure to keep `defer` — it works, but it is a function call per iteration to preserve a keyword you did not need.
- Assuming the reversed order is harmless; it usually is not for dependent resources.

# Deleting from a slice, order-preserving

## Intuition

Removing index `i` while keeping order joins the parts before and after it:

```go
return append(xs[:i], xs[i+1:]...)
```

The `...` spreads the tail as individual arguments.

## Approach

1. If i is out of [0, len(xs)), return xs unchanged.
2. Otherwise append(xs[:i], xs[i+1:]...) to splice out element i, preserving order.
3. Return the result.

## Solution

```go
func RemoveAt(xs []int, i int) []int {
	if i < 0 || i >= len(xs) {
		return xs
	}
	return append(xs[:i], xs[i+1:]...)
}
```

## Walkthrough

RemoveAt([1,2,3,4],1): xs[:1]=[1], spread xs[2:]=[3,4] -> [1,3,4].

## Pitfalls

- It **overwrites** `xs`'s backing array — copy first if the caller's slice
  matters.
- Guard `i` in `[0, len)` or the slice expressions panic.
- The leftover tail element lingers in the backing array (a leak for pointers).

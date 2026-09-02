# Pointer To Value

## Intuition

Go has no address-of for literals, but a parameter is an ordinary local variable. Returning its address is safe: the value outlives the call because the compiler moves it to the heap.

## Approach

1. Return `&v`.

## Solution

```go
func Ptr[T any](v T) *T {
	return &v
}
```

## Walkthrough

`Ptr(7)` copies `7` into the parameter `v`, then returns `&v`; dereferencing the result gives `7` back.

## Pitfalls

- Trying `return &7` — literals are not addressable.
- Expecting two calls with equal arguments to return the same pointer; each call allocates separately.
- Worrying about a dangling pointer — Go's escape analysis handles this.

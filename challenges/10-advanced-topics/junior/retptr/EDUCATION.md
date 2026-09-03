# A Pointer That Outlives Its Frame

## Intuition

Go has no dangling pointers because the compiler proves what escapes. Taking the address of a local and returning it does not break — it just means the local cannot live on the stack.

## Approach

1. Copy `v` into a local.
2. Return the local's address.

## Solution

```go
// New returns a pointer to a fresh int holding v.
//
// The pointer outlives the call, so the int cannot live in the frame — the
// compiler moves it to the heap. That is one allocation, and exactly one.
//
// Examples:
//
// 	*New(7) => 7
func New(v int) *int {
	p := v
	return &p
}
```

## Walkthrough

`p := v; return &p` — `p` is referenced after the frame dies, so the compiler allocates it on the heap. `new(int)` plus an assignment does the same thing.

## Pitfalls

- Returning `&v` directly works too; the parameter is a local like any other.
- Expecting zero allocations — a heap escape is the point here, not a bug.

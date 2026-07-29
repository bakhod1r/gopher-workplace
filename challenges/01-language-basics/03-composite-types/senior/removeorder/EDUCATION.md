# The delete idiom

## Intuition

Deleting index `i` keeps everything before it and everything **after** it:

```go
append(xs[:i], xs[i+1:]...)
```

Using `xs[i:]` for the tail re-includes the element you meant to drop.

## Approach

1. Bug: append(xs[:i], xs[i:]...) re-appends the element at index i (off-by-one), so nothing is removed. 2. Fix: skip the removed element with xs[i+1:]. 3. append(xs[:i], xs[i+1:]...) shifts the tail left over index i.

## Solution

```go
func RemoveAt(xs []int, i int) []int {
	return append(xs[:i], xs[i+1:]...)
}
```

## Walkthrough

xs=[1,2,3,4], i=1. Buggy xs[i:]=[2,3,4] -> result [1,2,3,4]. Correct xs[i+1:]=[3,4] -> [1,3,4].

## Pitfalls

- Tail starts at `i+1`, not `i`.
- This overwrites the backing array — copy first if the caller's slice matters.
- Guard `i` in range before slicing.

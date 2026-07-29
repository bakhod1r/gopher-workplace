# Arrays are value types

## Intuition

A Go array is a value: `b := *a` copies **all** elements. Mutating the copy
leaves the original array (reachable via the pointer) unchanged. Mutate through
the pointer directly:

```go
for i := range a { a[i] *= 2 } // a is *[3]int; indexing auto-derefs
```

## Approach

1. Bug: b := *a dereferences and copies the whole array by value; mutating b never touches the caller's array. 2. Fix: b := a keeps the pointer (*[3]int). 3. Ranging over a pointer-to-array and b[i] *= 2 auto-dereferences, mutating the original in place.

## Solution

```go
func Double(a *[3]int) {
	b := a
	for i := range b {
		b[i] *= 2
	}
}
```

## Walkthrough

*a copies [1,2,3] into b; doubling b leaves a=[1,2,3]. With b := a, b aliases a; b[i]*=2 writes through the pointer -> a=[2,4,6].

## Pitfalls

- `*a` (or `b := *a`) copies the array; index `a[i]` to mutate in place.
- Large arrays copied by value are costly — pass pointers or use slices.
- `[3]int` and `[4]int` are distinct types.

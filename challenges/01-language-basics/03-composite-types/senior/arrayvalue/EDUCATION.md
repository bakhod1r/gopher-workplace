# Arrays are value types

## The idea

A Go array is a value: `b := *a` copies **all** elements. Mutating the copy
leaves the original array (reachable via the pointer) unchanged. Mutate through
the pointer directly:

```go
for i := range a { a[i] *= 2 } // a is *[3]int; indexing auto-derefs
```

## Why it matters

The array-vs-slice distinction: assigning/copying an array duplicates its data,
while a slice shares a backing array. Accidentally working on a copy silently
no-ops the intended in-place update.

## Watch out

- `*a` (or `b := *a`) copies the array; index `a[i]` to mutate in place.
- Large arrays copied by value are costly — pass pointers or use slices.
- `[3]int` and `[4]int` are distinct types.

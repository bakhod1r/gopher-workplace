# Pointers to arrays

## Intuition

Arrays are values; a `*[N]T` lets a function mutate the caller's array, and indexing auto-dereferences the pointer.

## Approach

1. `arr` is a `*[4]int`.
2. `arr[i] *= 2` — Go auto-dereferences the array pointer for indexing, so the caller's array changes.

## Solution

```go
func Double(arr *[4]int, i int) {
	arr[i] *= 2
}
```

## Walkthrough

`Double(&a, 2)` with `a[2] = 3`: `arr[2] *= 2` doubles it to `6` in the caller's array.

## Pitfalls

- Passing `[4]int` by value copies; the caller wouldn't see changes.
- `arr[i]` on a pointer is `(*arr)[i]`.

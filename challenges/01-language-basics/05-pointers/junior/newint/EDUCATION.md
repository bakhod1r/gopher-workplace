# Returning pointers to locals

## Intuition

Unlike C, Go lets you return `&local`; escape analysis moves it to the heap so it stays valid.

## Approach

1. The parameter `v` already lives in fresh storage on each call.
2. Return its address with `&v`; Go moves it to the heap (escape analysis).
3. Every call gets its own `v`, so the pointers differ.

## Solution

```go
func Alloc(v int) *int {
	return &v
}
```

## Walkthrough

`Alloc(7)` binds `v = 7`, returns `&v`. Because `v` escapes, it survives the return; a second call has an independent `v`, so `p != q`.

## Pitfalls

- Returning `&v` is safe — no dangling pointer.
- `new(int)` also allocates a zeroed int and returns its pointer.

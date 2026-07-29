# Growing a slice via a slice pointer

## Intuition

Appending to a local dereference and dropping it leaves `*sp` stale; assign the result through the pointer.

## Approach

1. `s := *sp; s = append(...)` updates a local header only.
2. Write back through the pointer: `*sp = append(*sp, vs...)`.

## Solution

```go
func Extend(sp *[]int, vs ...int) {
	*sp = append(*sp, vs...)
}
```

## Walkthrough

The bug appends to a local copy of the header and discards it, so the caller's slice stays empty. Assigning through `*sp` publishes the grown slice.

## Pitfalls

- `s := *sp; s = append(...)` updates only the local.
- `*sp = append(*sp, vs...)` reaches the caller.

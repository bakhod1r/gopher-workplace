# Growing a slice via a slice pointer

## The idea

Appending to a local dereference and dropping it leaves `*sp` stale; assign the result through the pointer.

## Why it matters

Losing the append result behind a slice pointer is a silent no-op growth bug.

## Watch out

- `s := *sp; s = append(...)` updates only the local.
- `*sp = append(*sp, vs...)` reaches the caller.

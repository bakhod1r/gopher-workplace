# Swapping fields across instances

## The idea

Parallel assignment on two struct pointers exchanges a field without a temporary, mutating both callers' structs.

## Why it matters

In-place object edits (swap, rotate) work through struct pointers.

## Watch out

- Swapping the pointers (`a, b = b, a`) wouldn't affect the caller.
- Field-level swap touches the real instances.

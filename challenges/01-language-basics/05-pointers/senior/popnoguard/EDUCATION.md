# Guarding container operations

## The idea

Removing from an empty pointer-backed container must check for nil before dereferencing the head.

## Why it matters

Missing empty-guards are a common nil-panic in queues/stacks.

## Watch out

- `q.head.Val` on an empty queue panics.
- Return the zero/not-ok pair before dereferencing.

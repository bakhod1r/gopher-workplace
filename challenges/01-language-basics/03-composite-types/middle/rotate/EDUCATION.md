# Rotation by slice reassembly

## The idea

Rotating left by `k` moves the first `k` elements to the end. Normalize `k` into
`[0,n)` first (handles large and negative), then concatenate the two parts:

```go
k = ((k % n) + n) % n
out := append([]int{}, xs[k:]...)
out = append(out, xs[:k]...)
```

## Why it matters

Round-robin queues, ring buffers, and cyclic shifts all rotate. Normalizing the
count is the same modular idea as clock arithmetic.

## Watch out

- Guard `n == 0` before `k % n`.
- `append([]int{}, xs[k:]...)` copies, so the result doesn't alias `xs`.
- An in-place rotate (three reversals) is O(1) space but mutates the input.

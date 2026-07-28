# Range value copies

## The idea

The range loop's value variable is a per-iteration copy; mutating it never reaches the underlying element — only indexing does.

## Why it matters

Assuming `v` aliases the element is a frequent no-op mutation bug.

## Watch out

- `for _, v := range xs { v = ... }` changes nothing.
- Use `xs[i]` (or a pointer element) to mutate in place.

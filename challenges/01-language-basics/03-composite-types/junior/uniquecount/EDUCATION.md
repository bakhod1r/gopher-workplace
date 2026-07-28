# Sets from maps

## The idea

Go has no set type; you use a map whose keys are the members. The value carries
no information, so `struct{}` (zero bytes) is idiomatic:

```go
seen := make(map[int]struct{})
for _, x := range xs { seen[x] = struct{}{} }
return len(seen)
```

## Why it matters

Deduplication, membership, and "distinct count" are set operations. Inserting is
idempotent, and `len` gives the count for free.

## Watch out

- `map[int]struct{}` uses no memory for values; `map[int]bool` also works but
  stores a byte.
- Membership test is `_, ok := seen[x]`.
- The element type must be comparable to be a map key.

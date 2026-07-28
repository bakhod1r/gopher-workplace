# Comparing slices

## The idea

Slices are not comparable with `==` (except to `nil`). Equality is length plus
element-wise comparison:

```go
if len(a) != len(b) { return false }
for i := range a { if a[i] != b[i] { return false } }
return true
```

## Why it matters

A length-only check passes for same-size but different contents — a false
positive that corrupts caching, memoization, and change detection.

## Watch out

- `a == b` on slices is a compile error; don't reach for it.
- `slices.Equal` (Go 1.21+) does exactly this.
- `reflect.DeepEqual` also works but is slower and broader.

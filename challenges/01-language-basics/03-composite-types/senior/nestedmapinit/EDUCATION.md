# Nested maps need inner initialization

## The idea

`map[string]map[string]int` has an inner map per outer key. A missing outer key
reads as a **nil** inner map; writing `m[o][i]++` to nil panics. Lazily create it:

```go
if m[o] == nil { m[o] = make(map[string]int) }
m[o][i]++
```

## Why it matters

Multi-level maps (adjacency, counts-by-two-keys) are common. The outer `make`
doesn't create inner maps — each must be initialized before its first write.

## Watch out

- Reading `m[o][i]` is safe (nil inner returns zero); **writing** is not.
- The check-and-create must precede the write.
- Alternatives: a `map[[2]string]int` keyed by the pair avoids nesting.

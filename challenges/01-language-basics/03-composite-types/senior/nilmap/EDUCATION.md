# Nil maps: read yes, write no

## The idea

A `var m map[K]V` is nil. You can read it (missing keys return the zero value)
and range it (zero iterations), but **writing** to it panics:

```go
m := make(map[int]int) // must allocate before m[x]++
```

## Why it matters

`m[x]++` is a write. Functions that build and return a map must `make` it first.
The nil-map panic is a very common Go bug, often hidden until the first insert.

## Watch out

- Struct fields of map type are nil until you `make` them.
- Returning a nil map is fine if the caller only reads it — but tests/consumers
  may expect a non-nil empty map.
- Reading and `range` never panic on nil; only writes do.

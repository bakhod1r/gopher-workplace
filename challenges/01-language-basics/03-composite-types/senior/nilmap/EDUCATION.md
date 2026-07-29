# Nil maps: read yes, write no

## Intuition

A `var m map[K]V` is nil. You can read it (missing keys return the zero value)
and range it (zero iterations), but **writing** to it panics:

```go
m := make(map[int]int) // must allocate before m[x]++
```

## Approach

1. Bug: var m map[int]int leaves m nil; m[x]++ panics with assignment to entry in nil map. 2. A nil map is readable but not writable. 3. Fix: m := make(map[int]int).

## Solution

```go
func Count(xs []int) map[int]int {
	m := make(map[int]int)
	for _, x := range xs {
		m[x]++
	}
	return m
}
```

## Walkthrough

Ranging over [1,1,2] the first m[1]++ panics on a nil map. After make, m[1]++ -> 1, again -> 2, m[2]++ -> 1.

## Pitfalls

- Struct fields of map type are nil until you `make` them.
- Returning a nil map is fine if the caller only reads it — but tests/consumers
  may expect a non-nil empty map.
- Reading and `range` never panic on nil; only writes do.

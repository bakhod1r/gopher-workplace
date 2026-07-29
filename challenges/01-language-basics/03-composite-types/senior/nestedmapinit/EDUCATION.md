# Nested maps need inner initialization

## Intuition

`map[string]map[string]int` has an inner map per outer key. A missing outer key
reads as a **nil** inner map; writing `m[o][i]++` to nil panics. Lazily create it:

```go
if m[o] == nil { m[o] = make(map[string]int) }
m[o][i]++
```

## Approach

1. Bug: m[o][i]++ writes into m[o] which is a nil inner map when o is first seen -> panic (assignment to entry in nil map). 2. Fix: lazily initialize the inner map: if m[o]==nil { m[o]=make(map[string]int) }. 3. Then the ++ into the inner map is safe.

## Solution

```go
func Tally(pairs [][2]string) map[string]map[string]int {
	m := make(map[string]map[string]int)
	for _, p := range pairs {
		o, i := p[0], p[1]
		if m[o] == nil {
			m[o] = make(map[string]int)
		}
		m[o][i]++
	}
	return m
}
```

## Walkthrough

First pair (x,a): m[x] is nil, m[x][a]++ panics. After adding the init guard, m[x] becomes an empty map, then m[x][a] increments to 1.

## Pitfalls

- Reading `m[o][i]` is safe (nil inner returns zero); **writing** is not.
- The check-and-create must precede the write.
- Alternatives: a `map[[2]string]int` keyed by the pair avoids nesting.

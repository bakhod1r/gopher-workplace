# An Address Is Not A Reference

## Intuition

`unsafe.Pointer` is a pointer the garbage collector understands; `uintptr` is an integer that happens to look like one. Splitting the arithmetic across statements gives the runtime a window in which nothing refers to the object.

## Approach

1. Offset the pointer with `unsafe.Add(unsafe.Pointer(p), unsafe.Offsetof(p.B))`.
2. Convert the result to `*int64` and read through it.

## Solution

```go
import "unsafe"

// Pair is two 64-bit words.
type Pair struct {
	A int64
	B int64
}

// SecondWord returns the B field of the pair p points at, reached through
// the field's offset.
//
// Address arithmetic must stay in unsafe.Pointer. A uintptr is a plain
// number: nothing keeps the object alive and nothing updates it.
//
// Examples:
//
// 	SecondWord(&Pair{A: 1, B: 2}) => 2
func SecondWord(p *Pair) int64 {
	q := unsafe.Add(unsafe.Pointer(p), unsafe.Offsetof(p.B))
	return *(*int64)(q)
}
```

## Walkthrough

`unsafe.Add` produces a pointer into the same `Pair`, so the object stays reachable throughout. The `uintptr` version holds only a number between the two statements — correct today by luck, not by rule.

## Pitfalls

- Believing that Go's collector never moves objects; the rules are written so it may.
- Wrapping the broken version in `runtime.KeepAlive` — that addresses liveness, not the address going stale.

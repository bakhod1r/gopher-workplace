# Nil guards apply to unsafe pointers too

## Intuition

`unsafe.Pointer` can be nil; dereferencing it panics like any nil pointer, so the guard must come before the read.

## Approach

1. The bug dereferences `p` before the nil check → panic.
2. Check `p == nil` first, then read.

## Solution

```go
import "unsafe"

func ReadOr(p unsafe.Pointer, def int) int {
	if p == nil {
		return def
	}
	return *(*int)(p)
}
```

## Walkthrough

`ReadOr(nil, 7)` should return the default, but reading `*(*int)(p)` up front crashes. The nil guard must precede the load.

## Pitfalls

- Reorder: nil check, then dereference.
- unsafe doesn't exempt you from nil safety.

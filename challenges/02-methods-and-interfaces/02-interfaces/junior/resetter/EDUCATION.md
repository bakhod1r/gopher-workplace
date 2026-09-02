# Resetter

## Intuition

When an interface method mutates, the interface value must hold a pointer. Storing a struct value would make every `Reset` write to a copy.

## Approach

1. `(*Buffer).Reset` sets `b.Data = nil`.
2. `(*Gauge).Reset` sets `g.Value = 0`.
3. `ResetAll` ranges and calls `r.Reset()`.

## Solution

```go
func (b *Buffer) Reset() { b.Data = nil }

func (g *Gauge) Reset() { g.Value = 0 }

func ResetAll(rs []Resetter) {
	for _, r := range rs {
		r.Reset()
	}
}
```

## Walkthrough

`ResetAll([]Resetter{b, g})` holds two pointers. Each `Reset` writes through the pointer, so the caller's `b` and `g` are both cleared.

## Pitfalls

- Value receivers: the code compiles for direct calls but the caller's value is untouched — and `Buffer{}` would not satisfy `Resetter` in the slice.
- `b.Data = []string{}` also works, but `nil` is the idiomatic empty slice.
- Rebuilding the struct (`*b = Buffer{}`) is fine; reassigning the local `b` is not.

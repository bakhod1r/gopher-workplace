# Carve Aligned Blocks From One Allocation

## Intuition

An arena trades per-object freeing for near-zero allocation cost: a rounded-up cursor and a bounds check. Everything it hands out lives exactly as long as the arena does.

## Approach

1. Validate `n` and the power-of-two `align`.
2. Round `used` up to the alignment.
3. Reject the request if the block does not fit, comparing without overflowing.
4. Advance `used` and return the capped sub-slice.

## Solution

```go
// Arena hands out blocks of one fixed backing allocation.
type Arena struct {
	buf  []byte
	used uintptr
}

// NewArena returns an arena of size bytes.
func NewArena(size int) *Arena {
	if size < 0 {
		size = 0
	}
	return &Arena{buf: make([]byte, size)}
}

// Used reports how many bytes have been handed out, including padding.
func (a *Arena) Used() int { return int(a.used) }

// Cap reports the arena's total size.
func (a *Arena) Cap() int { return len(a.buf) }

// Alloc returns the next n bytes of the arena, starting at an offset that
// is a multiple of align.
//
// The arena never grows: when the remaining space cannot satisfy the
// request, Alloc reports false.
//
// Examples:
//
// 	a := NewArena(64); a.Alloc(8, 8) => an 8-byte block, true
func (a *Arena) Alloc(n int, align uintptr) ([]byte, bool) {
	if n < 0 || align == 0 || align&(align-1) != 0 {
		return nil, false
	}
	start := (a.used + align - 1) &^ (align - 1)
	if start > uintptr(len(a.buf)) || uintptr(n) > uintptr(len(a.buf))-start {
		return nil, false
	}
	a.used = start + uintptr(n)
	return a.buf[start : start+uintptr(n) : start+uintptr(n)], true
}
```

## Walkthrough

After a one-byte block, the cursor is 1; a request for 8 bytes with alignment 8 rounds to 8, hands out bytes 8..15, and leaves `used` at 16 — the seven skipped bytes are gone.

## Pitfalls

- `start + n > len(buf)` can overflow for a huge `n`; compare against the remaining space instead.
- Omitting the capacity cap, so an append to one block overwrites the next.
- Advancing `used` before deciding the request fits.

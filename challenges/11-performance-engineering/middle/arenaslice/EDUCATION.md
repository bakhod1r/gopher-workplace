# One Block, Many Objects

## Intuition

The block is a stretch of memory and `used` is how much of it is spoken for. Allocating moves the cursor forward; resetting moves it back to zero.

## Approach

1. `Alloc` bounds-checks, slices the window with a capped capacity, zeroes it, and advances the cursor.
2. `Reset` zeroes nothing eagerly — the zeroing happens when memory is handed out again — and sets `used` to 0.

## Solution

```go
func NewArena(n int) *Arena {
	if n < 0 {
		n = 0
	}
	return &Arena{block: make([]byte, n)}
}

func (a *Arena) Alloc(n int) ([]byte, bool) {
	if n <= 0 {
		return []byte{}, true
	}
	if n > a.Free() {
		return nil, false
	}
	out := a.block[a.used : a.used+n : a.used+n]
	clear(out)
	a.used += n
	return out, true
}

func (a *Arena) Used() int { return a.used }

func (a *Arena) Free() int { return len(a.block) - a.used }

func (a *Arena) Reset() { a.used = 0 }
```

## Walkthrough

The three-index slice `[lo:hi:hi]` sets the capacity equal to the length, so a caller who appends to their allocation gets a fresh array rather than silently overwriting the next object in the arena.

## Pitfalls

- Advancing `used` before the bounds check, so a failed allocation still consumes the block.
- Returning `a.block[a.used:]`, whose spare capacity reaches into everyone else's memory.
- Using a slice after `Reset` — the memory is still addressable, which is exactly what makes the bug hard to find.

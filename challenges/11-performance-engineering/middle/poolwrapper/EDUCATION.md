# A Typed Front Door For `sync.Pool`

## Intuition

A pool is a box of used objects. Whoever takes one out has to know it is used — so make the wrapper do the cleaning, once, in one place.

## Approach

1. `Get` pulls from the pool, falling back to a fresh `make` on a miss, and returns it with length zero.
2. `Put` ignores nil and stores the buffer back.

## Solution

```go
func (p *Pool) size() int {
	if p.Size <= 0 {
		return 1024
	}
	return p.Size
}

func (p *Pool) Get() []byte {
	if v := p.pool.Get(); v != nil {
		if b, ok := v.([]byte); ok && cap(b) >= p.size() {
			return b[:0]
		}
	}
	return make([]byte, 0, p.size())
}

func (p *Pool) Put(b []byte) {
	if b == nil {
		return
	}
	p.pool.Put(b[:0])
}
```

## Walkthrough

Resetting on both sides is deliberate: `Put` normalises what goes in, and `Get` normalises what comes out, so a buffer that reached the pool by another route still arrives empty.

## Pitfalls

- Returning the buffer without reslicing, so the next caller sees the previous one's bytes — a real information-leak class of bug.
- Asserting `v.([]byte)` without the comma-ok and panicking on an unexpected entry.
- Assuming the pool retains objects: a GC can empty it at any time, so `Get` must always be able to allocate.

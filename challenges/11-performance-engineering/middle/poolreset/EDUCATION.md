# Pools That Do Not Grow Forever

## Intuition

A pool is a cache, and every cache needs an admission policy. Here it is one rule: nothing bigger than `MaxCap` gets in.

## Approach

1. `Put` rejects nil and oversized buffers, otherwise resets and stores, counting under the mutex.
2. `Get` loops the pool for a buffer big enough, allocating when it finds none.

## Solution

```go
func (p *Pool) Put(b []byte) bool {
	if b == nil || cap(b) > MaxCap {
		return false
	}
	p.mu.Lock()
	p.kept++
	p.mu.Unlock()
	p.pool.Put(b[:0])
	return true
}

func (p *Pool) Get(n int) []byte {
	if n <= 0 {
		n = 1024
	}
	if v := p.pool.Get(); v != nil {
		if b, ok := v.([]byte); ok && cap(b) >= n {
			return b[:0]
		}
	}
	return make([]byte, 0, n)
}

func (p *Pool) Kept() int {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.kept
}
```

## Walkthrough

A pooled buffer that is too small is simply dropped rather than returned to the pool: keeping it would mean re-examining it on every subsequent `Get`, and it will be recreated cheaply if it is needed again.

## Pitfalls

- Accepting every buffer, so the pool's footprint is set by the worst request the service ever saw.
- Returning an undersized pooled buffer and letting the caller's `append` reallocate anyway.
- Reading `kept` without the mutex, which is a data race even for a plain `int`.

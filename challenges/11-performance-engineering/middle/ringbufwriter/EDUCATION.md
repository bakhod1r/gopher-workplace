# Bounded Memory For An Unbounded Stream

## Intuition

One array, one cursor. Writing past the end wraps to the front, and the cursor is simultaneously "where the next value goes" and "where the oldest value currently is".

## Approach

1. Store into `buf[next]`, advance `next` modulo the capacity, count the add.
2. Track the length separately until the ring fills.
3. `Snapshot` copies starting from the oldest position.

## Solution

```go
type Ring struct {
	buf   []int
	next  int
	count int64
	size  int
}

func New(n int) *Ring {
	if n < 0 {
		n = 0
	}
	return &Ring{buf: make([]int, n)}
}

func (r *Ring) Add(v int) {
	r.count++
	if len(r.buf) == 0 {
		return
	}
	r.buf[r.next] = v
	r.next = (r.next + 1) % len(r.buf)
	if r.size < len(r.buf) {
		r.size++
	}
}

func (r *Ring) Len() int { return r.size }

func (r *Ring) Total() int64 { return r.count }

func (r *Ring) Snapshot() []int {
	out := make([]int, 0, r.size)
	start := 0
	if r.size == len(r.buf) {
		start = r.next
	}
	for i := 0; i < r.size; i++ {
		out = append(out, r.buf[(start+i)%len(r.buf)])
	}
	return out
}
```

## Walkthrough

Before the ring fills, the oldest value is at index 0 and `next` is the length. Once it wraps, `next` points at the oldest entry — the slot about to be overwritten — which is where a chronological read must start.

## Pitfalls

- Reading from index 0 after wrapping, which reports the entries in the wrong order.
- Returning `r.buf` from `Snapshot`, letting the caller watch their data get overwritten.
- Growing the buffer when it fills, which is a queue, not a ring, and gives up the bounded memory.

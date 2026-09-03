# The Wraparound That Goes Negative

## Intuition

Go's `%` truncates toward zero, so a negative dividend yields a negative remainder. Stepping the head back from index 0 produces `-1`, which is not a valid slice index — the wrap must be forced positive by adding the buffer length first.

## Approach

1. Make room for the new element.
2. Step the head back by one, adding the buffer length before the remainder.
3. Store the value at the new head and grow the count.

## Solution

```go
func (d *Deque[T]) PushFront(v T) {
	d.grow()
	d.head = (d.head - 1 + len(d.buf)) % len(d.buf)
	d.buf[d.head] = v
	d.n++
}

func (d *Deque[T]) grow() {
	if d.n < len(d.buf) {
		return
	}
	nb := make([]T, 2*len(d.buf)+1)
	for i := 0; i < d.n; i++ {
		nb[i] = d.buf[(d.head+i)%len(d.buf)]
	}
	d.buf = nb
	d.head = 0
}

func (d *Deque[T]) PushBack(v T) {
	d.grow()
	d.buf[(d.head+d.n)%len(d.buf)] = v
	d.n++
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.n == 0 {
		var zero T
		return zero, false
	}
	v := d.buf[d.head]
	d.head = (d.head + 1) % len(d.buf)
	d.n--
	return v, true
}

func (d *Deque[T]) At(i int) (T, bool) {
	if i < 0 || i >= d.n {
		var zero T
		return zero, false
	}
	return d.buf[(d.head+i)%len(d.buf)], true
}

func (d *Deque[T]) Len() int {
	return d.n
}
```

## Walkthrough

With `buf` of length 3 and `head == 0`, `(0-1)%3` is `-1`, so `d.buf[-1]` panics immediately.

## Pitfalls

- Special-casing `head == 0` with an `if` — it works, but the arithmetic form is the one to know.
- Using `len(d.buf)` before `grow` has allocated anything, which divides by zero.
- Assuming `%` behaves like Python's, where the sign follows the divisor.

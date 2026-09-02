# Allocation-Bounded Sink

## Intuition

`append` amortises regrowth, but "amortised" still means the backing array is copied log(n) times and the old ones become garbage. When the final size is known, one `make` removes every copy.

## Approach

1. `NewSink` builds `make([]int, 0, n)` — length zero, capacity n.
2. `Write` appends; within capacity this is a store and a length bump, with no allocation.
3. `Len` returns `len(s.buf)`.
4. `Fill` loops `n` times calling `r.Write(i)`.

## Solution

```go
func NewSink(n int) *Sink {
	return &Sink{buf: make([]int, 0, n)}
}

func (s *Sink) Write(id int) { s.buf = append(s.buf, id) }

func (s *Sink) Len() int { return len(s.buf) }

func Fill(r Recorder, n int) {
	for i := 0; i < n; i++ {
		r.Write(i)
	}
}
```

## Walkthrough

`TestCapacityReserved` fills exactly to capacity and asserts `cap` never changed — the shape of the cost claim, made mechanical. `TestNoAllocsWithinCapacity` measures the same claim per call.

## Pitfalls

- `make([]int, n)` instead of `make([]int, 0, n)` — the sink starts with n zero records and `Len` is wrong.
- Storing the buffer by value in a value-receiver `Write`, so the length update is lost.
- Passing an `int` into an `any` parameter somewhere in the path, which boxes and allocates on every call.

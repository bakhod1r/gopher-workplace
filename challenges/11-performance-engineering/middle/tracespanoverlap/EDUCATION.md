# How Parallel Was It, Really

## Intuition

Forget the spans and keep only their edges. Walking the edges in time order with a counter tells you how many were open at every instant.

## Approach

1. Emit `+1` at each valid span's start and `−1` at its end.
2. Sort by time, ordering `−1` before `+1` when times are equal.
3. Sweep, tracking the running count, its maximum, and where the maximum first occurred.

## Solution

```go
func Intersect(a, b Span) (Span, bool) {
	start := max(a.Start, b.Start)
	end := min(a.End, b.End)
	if end <= start {
		return Span{}, false
	}
	return Span{start, end}, true
}

type event struct {
	t     int64
	delta int
}

func sweep(spans []Span) (peak int, at int64, ok bool) {
	events := make([]event, 0, 2*len(spans))
	for _, s := range spans {
		if s.End <= s.Start {
			continue
		}
		events = append(events, event{s.Start, +1}, event{s.End, -1})
	}
	if len(events) == 0 {
		return 0, 0, false
	}
	slices.SortFunc(events, func(x, y event) int {
		if c := cmp.Compare(x.t, y.t); c != 0 {
			return c
		}
		return cmp.Compare(x.delta, y.delta)
	})
	cur := 0
	for _, e := range events {
		cur += e.delta
		if cur > peak {
			peak, at = cur, e.t
		}
	}
	return peak, at, true
}

func Concurrency(spans []Span) int {
	peak, _, _ := sweep(spans)
	return peak
}

func BusiestAt(spans []Span) (int64, bool) {
	_, at, ok := sweep(spans)
	return at, ok
}
```

## Walkthrough

Sorting `−1` before `+1` at the same timestamp is what makes `{0,10}` and `{10,20}` report a concurrency of 1: the first span closes before the second opens, exactly as half-open intervals require.

## Pitfalls

- Ordering `+1` first on ties, which reports phantom overlap between adjacent spans.
- Updating `at` on `cur >= peak`, which reports the last peak instead of the first.
- Reading peak parallelism as CPU utilisation; a goroutine can be "running" in the trace while stalled on memory.

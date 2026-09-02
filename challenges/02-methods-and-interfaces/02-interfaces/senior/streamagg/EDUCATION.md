# Streaming Aggregate

## Intuition

An aggregator that keeps state proportional to the input is a memory bug waiting for a big day. Fold each reading into fixed-size state and the job runs in constant space no matter the volume.

## Approach

1. `MeanAgg.Add` updates `sum` and `count`; `Result` divides, guarding zero.
2. `MaxAgg` needs a `seenIt` flag so an all-negative stream is not reported as 0.
3. `Aggregate` loops `src.Next()` and calls `agg.Add`, returning `agg.Result()` at the end.
4. Nothing anywhere stores the readings.

## Solution

```go
func (m *MeanAgg) Add(v int) {
	m.sum += v
	m.count++
}

func (m *MeanAgg) Result() int {
	if m.count == 0 {
		return 0
	}
	return m.sum / m.count
}

func (m *MaxAgg) Add(v int) {
	if !m.seenIt || v > m.max {
		m.max = v
		m.seenIt = true
	}
}

func (m *MaxAgg) Result() int {
	if !m.seenIt {
		return 0
	}
	return m.max
}

func Aggregate(src Source, agg Aggregator) int {
	for {
		v, ok := src.Next()
		if !ok {
			return agg.Result()
		}
		agg.Add(v)
	}
}
```

## Walkthrough

Over 1M readings the allocation test passes only because nothing in the loop escapes: the source and the aggregator were allocated before the measured region, and `Add` writes to existing fields.

## Pitfalls

- Collecting readings into a slice and computing at the end — correct, and fatal at 100M rows.
- `MaxAgg` starting at zero without `seenIt`, which reports 0 for all-negative streams.
- Dividing before checking `count == 0`, which panics on an empty stream.

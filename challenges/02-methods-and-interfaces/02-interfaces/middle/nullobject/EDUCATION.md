# Null Object

## Intuition

A null object satisfies the interface and does nothing. Substituting it once at the boundary removes every downstream nil check, and the calls become unconditional.

## Approach

1. `(*Recorder).Report` appends the event.
2. `NopMetrics.Report` has an empty body.
3. `MetricsOr` returns `NopMetrics{}` for a nil interface, otherwise `m`.
4. `Process` reassigns `m = MetricsOr(m)` first, then reports in a plain loop.

## Solution

```go
func (r *Recorder) Report(name string) { r.Events = append(r.Events, name) }

func (NopMetrics) Report(name string) {}

func MetricsOr(m Metrics) Metrics {
	if m == nil {
		return NopMetrics{}
	}
	return m
}

func Process(m Metrics, items []string) int {
	m = MetricsOr(m)
	for _, item := range items {
		m.Report("item:" + item)
	}
	return len(items)
}
```

## Walkthrough

`Process(nil, []string{"a"})` replaces the nil interface with `NopMetrics{}`, so `m.Report` is a real call into an empty method — no panic, nothing recorded, count 1.

## Pitfalls

- Guarding every call with `if m != nil` — the pattern exists to delete those.
- Passing a typed nil (`var r *Recorder; Process(r, ...)`), which is *not* a nil interface and will panic inside `Report`.
- Returning `nil` from `MetricsOr` when `m` is nil, which defeats the whole exercise.

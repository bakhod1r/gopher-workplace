# Embedded State and Methods

## Intuition

When `Job` embeds `Tracker`, both fields (`Count`) and methods (`Inc`) are
promoted. You can interact with them directly on `j`.

## Approach

1. Call `j.Inc()`.

## Solution

```go
func (j *Job) Run() {
	j.Inc()
}
```

## Walkthrough

- `j.Run()` calls `j.Inc()`.
- `j.Inc()` is forwarded to `j.Tracker.Inc()`.
- `Tracker.Inc()` mutates `Tracker.Count`.
- `j.Count` reflects the change.

## Pitfalls

- Calling `j.Tracker.Inc()` is fine but less idiomatic.
- Modifying `j.Count++` manually bypasses the method, which is bad if `Inc`
  contained other logic (like locking).

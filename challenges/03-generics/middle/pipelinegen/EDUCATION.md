# Pipeline

## Intuition

Because every stage has the same type, the composition is a plain fold over the argument — no intermediate types to track, unlike `Compose`.

## Approach

1. Return a closure.
2. Apply each stage in order, reassigning `v`.
3. Return `v`.

## Solution

```go
func Pipeline[T any](stages ...func(T) T) func(T) T {
	return func(v T) T {
		for _, stage := range stages {
			v = stage(v)
		}
		return v
	}
}
```

## Walkthrough

`Pipeline(double, inc)(3)` computes `double(3) = 6`, then `inc(6) = 7`.

## Pitfalls

- Applying the stages in reverse.
- Ignoring a stage's result instead of reassigning `v`.
- Panicking or returning the zero value when there are no stages.

# Counter Deltas That Wrap Around

## Intuition

The constraint admits only unsigned types, so subtraction is modular: a decrease does not produce a negative delta, it produces a number close to the type's maximum. The reset has to be detected *before* the subtraction, not repaired afterwards.

## Approach

1. Emit the first sample as-is.
2. Compare each sample against its predecessor first.
3. Emit the sample itself on a reset, and the difference otherwise.

## Solution

```go
func Deltas[T Unsigned](samples []T) []T {
	out := make([]T, 0, len(samples))
	for i, v := range samples {
		if i == 0 || v < samples[i-1] {
			out = append(out, v)
			continue
		}
		out = append(out, v-samples[i-1])
	}
	return out
}
```

## Walkthrough

`Deltas([]uint8{10, 250, 5})` computes `5 - 250`, which is `11` in `uint8` — a plausible-looking number that is completely wrong.

## Pitfalls

- Fixing it with `if v-samples[i-1] < 0`, which is never true for an unsigned type.
- Widening to `int64` inside — it does not survive `uint64` inputs near the top of the range.

# The Flat Column

## Intuition

The profiler interrupts the program thousands of times and records which function was running. Counting those records per function *is* the flat column.

## Approach

1. Allocate the result map.
2. Range the samples, skipping non-positive values.
3. Accumulate into the map — the zero value of a missing key is already `0`.

## Solution

```go
func FlatSum(samples []Sample) map[string]int64 {
	out := make(map[string]int64)
	for _, s := range samples {
		if s.Self <= 0 {
			continue
		}
		out[s.Func] += s.Self
	}
	return out
}
```

## Walkthrough

`out[s.Func] += s.Self` needs no "does the key exist" check: reading a missing key yields `0`, so the first `+=` inserts.

## Pitfalls

- `var out map[string]int64`, which is nil and panics on write.
- Inserting the key before the guard, leaving `0`-valued entries.
- Confusing flat with cumulative and adding the callees' time too.

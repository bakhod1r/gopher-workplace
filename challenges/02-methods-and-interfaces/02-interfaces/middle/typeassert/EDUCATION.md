# Type Assertion

## Intuition

An assertion asks: is the dynamic type exactly this? The one-result form answers with a panic when it is not; the comma-ok form answers with a bool, which is what untrusted input needs.

## Approach

1. `AsInt` returns the result of `v.(int)` in comma-ok form.
2. `SumInts` ranges, calls `AsInt`, and adds only when `ok`.
3. Non-int payloads are skipped silently.

## Solution

```go
func AsInt(v any) (int, bool) {
	n, ok := v.(int)
	return n, ok
}

func SumInts(vs []any) int {
	sum := 0
	for _, v := range vs {
		if n, ok := AsInt(v); ok {
			sum += n
		}
	}
	return sum
}
```

## Walkthrough

`int64(5)` fails: assertions compare the dynamic type exactly, and `int64` is not `int`, so the result is `0, false` and `SumInts` skips it.

## Pitfalls

- `n := v.(int)` without `ok` — panics on the first non-int payload.
- Expecting numeric types to convert; assertion is not conversion.
- Returning `n` from a failed assertion as if it were meaningful — it is always the zero value.

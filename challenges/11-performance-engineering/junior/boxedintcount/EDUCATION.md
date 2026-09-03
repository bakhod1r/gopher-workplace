# What Boxing An `int` Costs

## Intuition

`any` stores a pointer, so the value has to exist somewhere addressable. For most ints that means a fresh heap word per conversion.

## Approach

1. Preallocate both results.
2. `Box` appends each value; the conversion is implicit.
3. `Unbox` uses the comma-ok assertion and counts the failures.

## Solution

```go
func Box(xs []int) []any {
	out := make([]any, 0, len(xs))
	for _, x := range xs {
		out = append(out, x)
	}
	return out
}

func Unbox(vs []any) ([]int, int) {
	out := make([]int, 0, len(vs))
	skipped := 0
	for _, v := range vs {
		if n, ok := v.(int); ok {
			out = append(out, n)
		} else {
			skipped++
		}
	}
	return out, skipped
}
```

## Walkthrough

Boxing 100 values in `0..255` costs one allocation — the slice — while boxing 100 values above 255 costs 101, because each one needs its own heap word.

## Pitfalls

- Benchmarking `[]any` code with small test values and concluding it is allocation-free.
- `v.(int)` without the comma-ok, which panics on the first string.
- Assuming `any` is free because it "just holds a value" — it holds a pointer to one.

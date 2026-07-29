# Conditional in-place mutation

## Intuition

Reading `*p`, testing bounds, and writing back constrains the caller's value without a return.

## Approach

1. If `*p < lo`, set `*p = lo`.
2. Else if `*p > hi`, set `*p = hi`.
3. Otherwise leave it alone.

## Solution

```go
func Clamp(p *int, lo, hi int) {
	if *p < lo {
		*p = lo
	} else if *p > hi {
		*p = hi
	}
}
```

## Walkthrough

`Clamp(&x, 0, 10)` with `x = 99`: `99 > 10`, so `*p = 10`.

## Pitfalls

- Only assign when out of range (optional micro-optimisation).
- Endpoints are inclusive.

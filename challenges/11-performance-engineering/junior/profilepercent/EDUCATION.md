# The Percentage Columns

## Intuition

A profile column is a ratio dressed up for humans: scale to 100, round to a fixed number of decimals, print.

## Approach

1. Guard a non-positive total.
2. Scale to a percentage, then round at the hundredths.
3. Format the rounded value with `%.2f`.

## Solution

```go
func Percent(value, total int64) float64 {
	if total <= 0 {
		return 0
	}
	return math.Round(float64(value)/float64(total)*100*100) / 100
}

func Format(value, total int64) string {
	return fmt.Sprintf("%.2f%%", Percent(value, total))
}
```

## Walkthrough

`1/3` is `33.3333...`; scaling by 100 twice gives `3333.33`, rounding gives `3333`, and dividing by 100 gives `33.33`.

## Pitfalls

- Integer division `value/total`, which is `0` for every proper fraction.
- Truncating instead of rounding, so `66.666` prints as `66.66`.
- Forgetting `%%` in the format string, which makes `Sprintf` eat the percent sign.

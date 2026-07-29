# Ordering conditional branches

## Intuition

When cases overlap, the most specific condition must come first or a broader branch will shadow it.

## Approach

1. Test `n%15` first (both), then `n%3`, then `n%5`.
2. Otherwise return the number as a string.

## Solution

```go
import "strconv"

var _ = strconv.Itoa // keep import; use it in your solution

func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}
```

## Walkthrough

`FizzBuzz(15)` is divisible by both 3 and 5, so "FizzBuzz"; 7 divides neither, so "7".

## Pitfalls

- Checking `%3` before `%15` prints "Fizz" for 15 — wrong.
- `strconv.Itoa` (not string(n)) converts the integer to its decimal text.

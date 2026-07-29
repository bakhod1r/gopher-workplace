# Money as integer minor units

## Intuition

Represent currency as an integer number of cents, never a float. Parse the
dollars and the (optional) 1–2 fraction digits separately, pad the fraction to
two places, and combine: `dollars*100 + cents`.

## Approach

1. Empty -> (0,false). 2. Scan digits into dollars until a '.' or end. 3. If a '.' follows, the fraction must be 1 or 2 digits (else false). 4. A single fractional digit is scaled *10. 5. All chars must be digits. 6. Return dollars*100 + cents.

## Solution

```go
func Cents(s string) (int, bool) {
	if s == "" {
		return 0, false
	}
	i := 0
	dollars := 0
	sawDigit := false
	for ; i < len(s) && s[i] != '.'; i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, false
		}
		dollars = dollars*10 + int(c-'0')
		sawDigit = true
	}
	cents := 0
	if i < len(s) {
		i++
		frac := s[i:]
		if len(frac) == 0 || len(frac) > 2 {
			return 0, false
		}
		for j := 0; j < len(frac); j++ {
			c := frac[j]
			if c < '0' || c > '9' {
				return 0, false
			}
			cents = cents*10 + int(c-'0')
		}
		if len(frac) == 1 {
			cents *= 10
		}
		sawDigit = true
	}
	if !sawDigit {
		return 0, false
	}
	return dollars*100 + cents, true
}
```

## Walkthrough

Cents("3.5"): dollars=3, frac="5" len 1 -> cents=5*10=50. 3*100+50=350 -> (350,true).

## Pitfalls

- `"3.5"` means 350 cents — pad a single fraction digit to two.
- Reject 3+ fraction digits rather than silently rounding.
- No `.` at all is valid ("7" → 700); a trailing `.` is not.

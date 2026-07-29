# Unit conversion factors

## Intuition

Parse a number, then multiply by its unit's factor in the common base (seconds):
hours ×3600, minutes ×60, seconds ×1. A wrong factor silently rescales every
value that uses that unit.

## Approach

1. Bug: minutes used `num * 6` instead of `num * 60`.
2. Fix: multiply minutes by 60 seconds.
3. Hours (3600) and seconds (1) were already correct.

## Solution

```go
func Seconds(s string) (int, bool) {
	total, num, seen := 0, 0, false
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= '0' && c <= '9':
			num = num*10 + int(c-'0')
			seen = true
		case c == 'h':
			if !seen {
				return 0, false
			}
			total += num * 3600
			num, seen = 0, false
		case c == 'm':
			if !seen {
				return 0, false
			}
			total += num * 60
			num, seen = 0, false
		case c == 's':
			if !seen {
				return 0, false
			}
			total += num
			num, seen = 0, false
		default:
			return 0, false
		}
	}
	if seen {
		return 0, false // trailing number without a unit
	}
	return total, true
}
```

## Walkthrough

"1h30m": 1*3600=3600, 30*60=1800, total 5400.

## Pitfalls

- Keep the factors exact: 60 seconds per minute, 3600 per hour.
- Require a unit after each number; a dangling number is malformed.
- Empty input is a valid zero here — decide such edge cases explicitly.

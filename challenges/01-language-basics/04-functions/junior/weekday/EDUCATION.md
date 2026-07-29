# expression switch

## Intuition

Switching on a value selects the matching case; `default` catches unmatched inputs cleanly.

## Approach

1. Switch on `d` with a case per day 1–7.
2. `default` returns "Unknown".

## Solution

```go
func Weekday(d int) string {
	switch d {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Unknown"
	}
}
```

## Walkthrough

`Weekday(1)` matches case 1 → "Monday"; 8 has no case, so "Unknown".

## Pitfalls

- Go cases don't fall through by default.
- `default` handles 0, 8, negatives.

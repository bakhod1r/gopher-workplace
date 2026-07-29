# Value switch for lookups

## Intuition

`switch x { case a: }` compares `x` against each case value and runs the first equal one; `default` handles the rest.

## Approach

1. `switch d` over 0..6.
2. `default` returns "?" for out-of-range.

## Solution

```go
func DayName(d int) string {
	switch d {
	case 0:
		return "Sun"
	case 1:
		return "Mon"
	case 2:
		return "Tue"
	case 3:
		return "Wed"
	case 4:
		return "Thu"
	case 5:
		return "Fri"
	case 6:
		return "Sat"
	default:
		return "?"
	}
}
```

## Walkthrough

`DayName(0)` is "Sun"; 9 has no case, so the default "?".

## Pitfalls

- Provide a `default` so unexpected input has a defined result.
- Cases can group values: `case 0, 6:` for weekend logic.

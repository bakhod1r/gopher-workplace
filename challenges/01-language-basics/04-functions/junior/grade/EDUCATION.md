# if/else-if ladder

## Intuition

Ordered threshold checks classify a value; testing highest first keeps each band correct.

## Approach

1. Test the highest threshold first.
2. Use `>=` for inclusive boundaries.
3. Fall to "F" below 60.

## Solution

```go
func Grade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}
```

## Walkthrough

`Grade(70)` fails `>= 90` and `>= 80` but passes `>= 70` → "C".

## Pitfalls

- Order matters: a low threshold first would swallow higher scores.
- Use `>=` for inclusive boundaries.

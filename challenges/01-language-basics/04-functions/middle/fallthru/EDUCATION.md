# The fallthrough keyword

## Intuition

Unlike C, Go breaks after each case; `fallthrough` opts into entering the next case body (without checking its condition).

## Approach

1. Use `fallthrough` to accumulate lower tiers.
2. Join reached tiers with `/`.

## Solution

```go
func Rank(score int) string {
	result := ""
	add := func(s string) {
		if result != "" {
			result += "/"
		}
		result += s
	}
	switch {
	case score >= 9:
		add("gold")
		fallthrough
	case score >= 6:
		add("silver")
		fallthrough
	case score >= 3:
		add("bronze")
	}
	return result
}
```

## Walkthrough

`Rank(9)` enters the gold case and falls through silver and bronze, joining all three.

## Pitfalls

- `fallthrough` enters the NEXT case unconditionally — order your cases so that's correct.
- It must be the last statement in a case.

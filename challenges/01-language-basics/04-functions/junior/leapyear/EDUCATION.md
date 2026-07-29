# Composing layered rules

## Intuition

Nested exceptions collapse into a single boolean expression when you order the divisibility tests by precedence.

## Approach

1. Divisible by 400 → leap.
2. Divisible by 4 but not 100 → leap.
3. Otherwise not.

## Solution

```go
func IsLeap(y int) bool {
	return y%400 == 0 || (y%4 == 0 && y%100 != 0)
}
```

## Walkthrough

2000 is divisible by 400 → leap; 1900 is divisible by 100 but not 400 → not leap.

## Pitfalls

- 1900 is divisible by 4 and 100 but not 400, so it is NOT a leap year.
- Parenthesise to keep `&&` binding tighter than `||`.

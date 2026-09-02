# Stringer and Formatting

## Intuition

`fmt` checks every operand for `String() string` at runtime. Implementing it once changes every print site — logs, errors, `%v` in tests.

## Approach

1. Handle the negative case by recursing on the positive value and prefixing `-`.
2. Format with `%d.%02d` over `n/100` and `n%100`.
3. `Level.String` switches with a `default` that formats the raw number.
4. `Line` uses `%v` for both Stringers.

## Solution

```go
func (m Money) String() string {
	n := int(m)
	if n < 0 {
		return "-" + Money(-n).String()
	}
	return fmt.Sprintf("%d.%02d", n/100, n%100)
}

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Error:
		return "ERROR"
	default:
		return fmt.Sprintf("LEVEL(%d)", int(l))
	}
}

func Line(l Level, msg string, m Money) string {
	return fmt.Sprintf("[%v] %s: %v", l, msg, m)
}
```

## Walkthrough

`Money(-250)`: the negative branch recurses on 250, producing `"2.50"`, and prefixes the sign — `"-2.50"`. Formatting `-250/100` directly would give `-2` and `-50`, printing `"-2.-50"`.

## Pitfalls

- Using `%d.%d` — `Money(5)` then prints `"0.5"` instead of `"0.05"`.
- Calling `fmt.Sprintf("%v", l)` inside `Level.String` — infinite recursion.
- Letting `n % 100` run on a negative number and emitting a second minus sign.

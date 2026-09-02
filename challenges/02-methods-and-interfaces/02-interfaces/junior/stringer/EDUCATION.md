# Stringer

## Intuition

`fmt` checks at runtime whether a value implements `String() string`. Implementing it changes how the value prints everywhere, with no call-site changes.

## Approach

1. `Color.String` uses a `switch` with a `default` returning `"unknown"`.
2. `Temp.String` converts with `strconv.Itoa(int(t))` and appends `"C"`.
3. `Print` returns `s.String()`.

## Solution

```go
func (c Color) String() string {
	switch c {
	case Red:
		return "red"
	case Green:
		return "green"
	case Blue:
		return "blue"
	default:
		return "unknown"
	}
}

func (t Temp) String() string { return strconv.Itoa(int(t)) + "C" }

func Print(s Stringer) string { return s.String() }
```

## Walkthrough

`fmt.Sprintf("%v", Green)` finds `Color.String` through the `fmt.Stringer` check and prints `"green"` instead of `1`.

## Pitfalls

- Calling `fmt.Sprintf("%v", c)` inside `Color.String` — infinite recursion.
- `string(t)` instead of `strconv.Itoa(int(t))` — that builds a rune, not digits.
- Forgetting the `default` case, so unknown codes fall through to the zero value.

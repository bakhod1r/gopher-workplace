# Explicit fallthrough in Go

## Intuition

Unlike C, Go breaks after each case; `fallthrough` opts in and ignores the next case's condition — misuse silently misroutes values.

## Approach

1. Go `switch` cases do **not** fall through by default.
2. The stray `fallthrough` after `low` leaks into the next case, mislabeling.
3. Remove it.

## Solution

```go
func Label(level int) string {
	label := "?"
	switch level {
	case 1:
		label = "low"
	case 2:
		label = "mid"
	case 3:
		label = "high"
	}
	return label
}
```

## Walkthrough

With `fallthrough`, `Label(1)` continues into the `mid` branch and returns the wrong label. Dropping it makes each case independent.

## Pitfalls

- Only add `fallthrough` when you truly want the next case body.
- Code after an unconditional `return` is dead.

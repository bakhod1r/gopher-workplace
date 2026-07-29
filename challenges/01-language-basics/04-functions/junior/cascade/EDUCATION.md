# switch fallthrough

## Intuition

Unlike C, Go `switch` cases stop by default; `fallthrough` opts into continuing so lower tiers accumulate.

## Approach

1. Switch on the exact `level` (3, 2, 1).
2. Each case adds its permission then `fallthrough` to the lower one.
3. Unmatched levels return "".

## Solution

```go
func Access(level int) string {
	result := ""
	add := func(s string) {
		if result != "" {
			result += ","
		}
		result += s
	}
	switch level {
	case 3:
		add("admin")
		fallthrough
	case 2:
		add("write")
		fallthrough
	case 1:
		add("read")
	}
	return result
}
```

## Walkthrough

`Access(3)` matches case 3, adds "admin", falls through to add "write" and "read". `Access(9)` matches no case and returns empty.

## Pitfalls

- `fallthrough` transfers to the next case unconditionally, ignoring its condition.
- The last case cannot `fallthrough`.

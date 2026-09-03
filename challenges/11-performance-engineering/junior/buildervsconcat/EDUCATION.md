# Building A String Without The Garbage

## Intuition

You cannot append to a string, only build a new one. Doing that once, into a buffer of exactly the right size, is the whole optimisation.

## Approach

1. Return `""` for no parts.
2. Sum the part lengths plus `len(sep) * (len(parts)-1)`.
3. `Grow` by that amount and write everything.

## Solution

```go
func JoinSep(parts []string, sep string) string {
	if len(parts) == 0 {
		return ""
	}
	n := len(sep) * (len(parts) - 1)
	for _, p := range parts {
		n += len(p)
	}
	var b strings.Builder
	b.Grow(n)
	for i, p := range parts {
		if i > 0 {
			b.WriteString(sep)
		}
		b.WriteString(p)
	}
	return b.String()
}
```

## Walkthrough

`Builder.String` hands back the buffer without copying it, so the single `Grow` allocation is the only one in the whole function.

## Pitfalls

- `len(sep) * len(parts)`, which over-allocates by one separator.
- Writing the separator after each part and trimming the tail — correct output, extra work.
- Reusing a `Builder` after `String()` without `Reset`, which appends to the string you just returned.

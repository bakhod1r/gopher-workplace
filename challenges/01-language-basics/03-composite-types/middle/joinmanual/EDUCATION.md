# Joining strings

## Intuition

Write the separator **before** each element except the first, building into a
`strings.Builder`:

```go
var b strings.Builder
for i, p := range parts {
	if i > 0 { b.WriteString(sep) }
	b.WriteString(p)
}
return b.String()
```

## Approach

1. Use a strings.Builder.
2. Iterate parts with index.
3. Write sep before every part except the first.
4. Write the part.
5. Return builder string.

## Solution

```go
import "strings"

func Join(parts []string, sep string) string {
	var b strings.Builder
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

["a","b","c"]: i0 write "a"; i1 write "," then "b"; i2 write "," then "c" -> "a,b,c".

## Pitfalls

- Don't put the separator after the last element.
- Use `strings.Builder`, not repeated `+=`, for many parts.
- Empty slice yields "".

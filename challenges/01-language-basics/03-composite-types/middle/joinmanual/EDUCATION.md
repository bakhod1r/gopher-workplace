# Joining strings

## The idea

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

## Why it matters

Naive `result += sep + p` in a loop is O(n²) because strings are immutable —
every `+=` copies. `strings.Builder` amortizes to O(n). It also shows the
"separator between, not around" pattern.

## Watch out

- Don't put the separator after the last element.
- Use `strings.Builder`, not repeated `+=`, for many parts.
- Empty slice yields "".

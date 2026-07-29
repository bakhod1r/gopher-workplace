# CSV escaping

## Intuition

RFC 4180 says a field is quoted only when it contains a delimiter (comma), a
quote, or a newline. When quoted, every inner `"` is written as `""`:

```go
if strings.ContainsAny(s, ",\"\n") {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
return s
```

## Approach

1. Scan for comma, double-quote, or newline. 2. If none present, return s unchanged. 3. Otherwise emit a leading '"', copy each char doubling any '"', then a trailing '"'.

## Solution

```go
func Quote(s string) string {
	need := false
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ',', '"', '\n':
			need = true
		}
	}
	if !need {
		return s
	}
	out := []byte{'"'}
	for i := 0; i < len(s); i++ {
		if s[i] == '"' {
			out = append(out, '"', '"')
		} else {
			out = append(out, s[i])
		}
	}
	out = append(out, '"')
	return string(out)
}
```

## Walkthrough

Quote("a,b"): comma found -> wrap: '"' + a,b + '"' = "a,b" quoted.

## Pitfalls

- Doubling quotes must happen *before* wrapping, on the content only.
- A plain field is returned unchanged — don't over-quote.
- The real `encoding/csv` writer does this; here you learn the rule directly.

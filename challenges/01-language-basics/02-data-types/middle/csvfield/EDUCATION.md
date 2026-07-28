# CSV escaping

## The idea

RFC 4180 says a field is quoted only when it contains a delimiter (comma), a
quote, or a newline. When quoted, every inner `"` is written as `""`:

```go
if strings.ContainsAny(s, ",\"\n") {
	return `"` + strings.ReplaceAll(s, `"`, `""`) + `"`
}
return s
```

## Why it matters

Naive CSV writers that always quote, or never escape, produce files that break
downstream parsers. Getting the conditional-quote-and-double-quote rule right is
real, load-bearing correctness.

## Watch out

- Doubling quotes must happen *before* wrapping, on the content only.
- A plain field is returned unchanged — don't over-quote.
- The real `encoding/csv` writer does this; here you learn the rule directly.

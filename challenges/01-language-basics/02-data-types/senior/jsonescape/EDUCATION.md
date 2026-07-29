# Escaping the escape character

## Intuition

In JSON, the backslash introduces every escape, so a literal backslash must
itself be written as `\\`. If you escape quotes and newlines but forget the
backslash, any input containing `\` yields invalid JSON:

```go
case '\\':
	b.WriteString(`\\`) // raw literal: two backslash characters
```

## Approach

1. Bug: the backslash case was missing, so `\` passed through unescaped (producing invalid JSON).
2. Fix: add `case '\\': b.WriteString("\\\\")` before the quote case.
3. Quote, newline, tab, CR cases stay.

## Solution

```go
import "strings"

func Escape(s string) string {
	var b strings.Builder
	for _, r := range s {
		switch r {
		case '\\':
			b.WriteString(`\\`)
		case '"':
			b.WriteString(`\"`)
		case '\n':
			b.WriteString(`\n`)
		case '\t':
			b.WriteString(`\t`)
		case '\r':
			b.WriteString(`\r`)
		default:
			b.WriteRune(r)
		}
	}
	return b.String()
}
```

## Walkthrough

a\b: the \ now matches the backslash case and emits \\ -> a\\b.

## Pitfalls

- Escape the backslash *before* other rules conceptually; order in a `switch`
  doesn't matter since each rune matches once.
- A raw literal `` `\\` `` is two characters; `"\\\\"` is the interpreted
  equivalent.
- Prefer `encoding/json` in production; hand-escaping is for understanding.

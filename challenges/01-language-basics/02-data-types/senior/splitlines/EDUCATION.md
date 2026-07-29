# Normalizing line endings

## Intuition

Split on `\n`, then strip a trailing `\r` from each piece so CRLF and LF inputs
produce identical lines:

```go
parts[i] = strings.TrimSuffix(p, "\r")
```

## Approach

1. Bug: `parts[i] = p` left the trailing '\r' on CRLF lines.
2. Fix: `parts[i] = strings.TrimSuffix(p, "\r")`.
3. LF and CRLF now yield identical results.

## Solution

```go
import "strings"

func Lines(s string) []string {
	parts := strings.Split(s, "\n")
	for i, p := range parts {
		parts[i] = strings.TrimSuffix(p, "\r")
	}
	return parts
}
```

## Walkthrough

"a\r\nb": Split on \n -> ["a\r","b"]; TrimSuffix -> ["a","b"].

## Pitfalls

- `TrimSuffix` removes `\r` only if present — safe for LF-only input.
- Don't `TrimSpace`; that would also drop meaningful leading/trailing spaces.
- Some files end with a trailing newline, yielding a final empty line — decide
  whether to keep it.

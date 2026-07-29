# Normalizing text to a slug

## Intuition

Scan runes, keep alphanumerics (lowercased), and collapse every run of other
characters into a single `-`. A neat trick avoids trailing dashes: instead of
appending a dash on each separator, set a *pending* flag and only emit the dash
right before the next kept character.

## Approach

1. Scan bytes. 2. Uppercase A-Z lowered and appended; a-z/0-9 appended as-is. 3. Any other char emits a single '-' only if output is non-empty and the previous char wasn't already a dash (collapses runs, drops leading). 4. Trim trailing '-'.

## Solution

```go
func Slug(s string) string {
	var out []byte
	prevDash := false
	for i := 0; i < len(s); i++ {
		c := s[i]
		switch {
		case c >= 'A' && c <= 'Z':
			out = append(out, c+('a'-'A'))
			prevDash = false
		case (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9'):
			out = append(out, c)
			prevDash = false
		default:
			if len(out) > 0 && !prevDash {
				out = append(out, '-')
				prevDash = true
			}
		}
	}
	for len(out) > 0 && out[len(out)-1] == '-' {
		out = out[:len(out)-1]
	}
	return string(out)
}
```

## Walkthrough

Slug("a---b"): 'a'->a; '-'x3 -> one '-' (prevDash guards rest); 'b'->b -> "a-b".

## Pitfalls

- Lowercase with `unicode.ToLower` (or ASCII arithmetic).
- Don't emit a leading dash: the pending flag naturally handles this if you only
  flush before real output.
- Decide how to treat non-ASCII letters (this spec keeps only ASCII alnum).

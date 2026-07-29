# Bounded character ranges

## Intuition

ASCII uppercase letters occupy 65..90. Lowercasing adds 32 — but only for that
exact range. A one-sided check (`c >= 'A'`) also matches `[`, `\`, `]`, `^`, `_`,
backtick, and every lowercase letter, corrupting them:

```go
if c >= 'A' && c <= 'Z' { b[i] = c + 32 }
```

## Approach

1. Bug: condition `c >= 'A'` also matched bytes above 'Z' ([, \, ], _, and a-z), corrupting them.
2. Fix: `c >= 'A' && c <= 'Z'` restricts to uppercase letters.
3. Only A-Z get +32.

## Solution

```go
func Lower(s string) string {
	b := []byte(s)
	for i, c := range b {
		if c >= 'A' && c <= 'Z' {
			b[i] = c + 32
		}
	}
	return string(b)
}
```

## Walkthrough

"a[b]": '['=91 > 'Z'=90, now excluded, stays '[' -> "a[b]".

## Pitfalls

- Always bound both ends of a character range.
- `unicode.ToLower` handles non-ASCII correctly; this ASCII-only version is for
  bytes.
- The `+32` offset is specific to ASCII letters.

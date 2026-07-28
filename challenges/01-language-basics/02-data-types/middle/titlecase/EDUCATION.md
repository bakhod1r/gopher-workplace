# Case conversion

## The idea

ASCII upper/lower differ by one bit (0x20). Use `unicode.ToUpper` / `ToLower` for
clarity, and track whether you are at the start of a word:

```go
start := true
for _, r := range s {
	if r == ' ' { start = true; ...append space; continue }
	if start { append ToUpper(r) } else { append ToLower(r) }
	start = false
}
```

## Why it matters

It combines rune scanning, case conversion, and stateful iteration — a common
text-normalization shape.

## Watch out

- `strings.Title` is deprecated (Unicode-incorrect for many scripts); this ASCII
  exercise avoids that dependency.
- Build output with `strings.Builder` or `[]rune`, not by `+=` in a loop.
- Multiple/leading spaces change word boundaries — decide the spec.

# Truncating UTF-8 safely

## The idea

`for i, r := range s` gives the **byte index** where each rune starts. A rune
occupies `utf8.RuneLen(r)` bytes. Keep runes while the next one still fits under
the byte budget; stop at the boundary:

```go
for i, r := range s {
	if i+utf8.RuneLen(r) > max { return s[:i] }
}
return s
```

## Why it matters

Byte-counted limits (DB columns, wire fields, filesystem names) are everywhere.
Slicing `s[:max]` blindly can split a character and emit invalid UTF-8; cutting
on a rune boundary is the correct, real-world fix.

## Watch out

- The range index is a byte offset, exactly the safe cut point.
- `s[:i]` slices bytes without copying — cheap.
- `max` larger than `len(s)` returns the whole string; `max==0` returns "".

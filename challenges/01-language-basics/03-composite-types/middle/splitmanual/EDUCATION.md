# Splitting a string

## The idea

Remember the start of the current field; when you hit a separator, emit
`s[start:i]` and move `start` past it. After the loop, emit the final field:

```go
start := 0
for i := 0; i < len(s); i++ {
	if s[i] == sep { out = append(out, s[start:i]); start = i + 1 }
}
out = append(out, s[start:])
```

## Why it matters

It reveals what `Split` guarantees: N separators produce N+1 fields, including
empties. Slicing the string is cheap — the fields share the source's bytes.

## Watch out

- The result always has `count(sep)+1` elements — empty string gives `[""]`.
- Fields are sub-strings sharing memory; no copying.
- This splits on a byte; a multi-byte or rune separator needs more care.

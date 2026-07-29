# Splitting a string

## Intuition

Remember the start of the current field; when you hit a separator, emit
`s[start:i]` and move `start` past it. After the loop, emit the final field:

```go
start := 0
for i := 0; i < len(s); i++ {
	if s[i] == sep { out = append(out, s[start:i]); start = i + 1 }
}
out = append(out, s[start:])
```

## Approach

1. Track start index.
2. Scan; at each sep append s[start:i] and set start=i+1.
3. After the loop append the final segment s[start:].
4. Result has seps+1 elements.

## Solution

```go
func Split(s string, sep byte) []string {
	out := []string{}
	start := 0
	for i := 0; i < len(s); i++ {
		if s[i] == sep {
			out = append(out, s[start:i])
			start = i + 1
		}
	}
	out = append(out, s[start:])
	return out
}
```

## Walkthrough

"a,,c": i1 sep append "a" start=2; i2 sep append "" (s[2:2]) start=3; end append "c" -> ["a","","c"].

## Pitfalls

- The result always has `count(sep)+1` elements — empty string gives `[""]`.
- Fields are sub-strings sharing memory; no copying.
- This splits on a byte; a multi-byte or rune separator needs more care.

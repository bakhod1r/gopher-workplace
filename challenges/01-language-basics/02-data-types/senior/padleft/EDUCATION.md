# Padding side and width unit

## Intuition

Left-padding **prepends** fill so the original text sits at the right edge:

```go
return pad + s
```

The pad count is `width - len([]rune(s))`, measured in runes so alignment is
visual, not byte-based.

## Approach

1. Bug: `return s + pad` appended the fill (right-pad).
2. Fix: `return pad + s` to left-pad.
3. Strings already >= width return unchanged.

## Solution

```go
import "strings"

func Pad(s string, width int, fill byte) string {
	n := width - len([]rune(s))
	if n <= 0 {
		return s
	}
	pad := strings.Repeat(string(fill), n)
	return pad + s
}
```

## Walkthrough

Pad("42",5,'0'): n=3, pad="000", pad+s -> "00042".

## Pitfalls

- Measure width in runes; byte width misaligns multi-byte text.
- Return unchanged when already at/over width — never truncate here.
- `strings.Repeat` builds the fill run efficiently.

# Inclusive width boundaries

## Intuition

A word fits on the current line when the resulting length — existing text, one
space, and the word — is **at most** the width. The boundary is inclusive:

```go
if len(line)+1+len(w) <= width { line += " " + w }
```

Using `<` rejects a line that exactly fills the width, wrapping one character too
soon.

## Approach

1. Bug: off-by-one `len(line)+1+len(w) < width` broke a line that exactly fit the width.
2. Fix: use `<= width` so a line reaching exactly width is allowed.
3. Words that would overflow start a new line.

## Solution

```go
import "strings"

func Wrap(text string, width int) []string {
	words := strings.Fields(text)
	var lines []string
	var line string
	for _, w := range words {
		if line == "" {
			line = w
			continue
		}
		if len(line)+1+len(w) <= width {
			line += " " + w
		} else {
			lines = append(lines, line)
			line = w
		}
	}
	if line != "" {
		lines = append(lines, line)
	}
	return lines
}
```

## Walkthrough

"aa bb cc" w5: "aa"+ " bb" len5 <=5 ok -> "aa bb"; adding " cc" len8 >5 -> new line "cc".

## Pitfalls

- Include the separator (`+1`) in the length math.
- `<=` vs `<` is the whole bug — inclusive width.
- This simplified version counts bytes; real wrapping counts runes/display width.

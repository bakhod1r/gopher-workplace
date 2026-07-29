# Split vs Fields

## Intuition

`strings.Split(s, " ")` cuts on **every** single space, so consecutive spaces
produce empty strings. `strings.Fields(s)` splits on runs of *any* whitespace and
drops the empties:

```go
len(strings.Fields(s)) // correct word count
```

## Approach

1. Bug: strings.Split(s," ") emits empty tokens for adjacent/leading/trailing spaces, inflating the count.
2. Fix: use strings.Fields, which splits on runs of whitespace and drops empties.
3. len(Fields) is the true word count.

## Solution

```go
import "strings"

func Count(s string) int {
	return len(strings.Fields(s))
}
```

## Walkthrough

"  a   b  " -> Fields -> ["a","b"] -> len 2. Split would give many "" tokens.

## Pitfalls

- `Fields` treats tabs and newlines as separators too — usually what you want.
- `Split("", " ")` returns `[""]` (length 1); `Fields("")` returns length 0.
- If you need a custom separator set, `strings.FieldsFunc` takes a predicate.

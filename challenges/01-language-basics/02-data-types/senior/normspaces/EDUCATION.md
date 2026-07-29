# Fields + Join normalization

## Intuition

`strings.Fields(s)` splits on runs of any whitespace and drops empties;
`strings.Join(..., " ")` rejoins with single spaces. Together they collapse and
trim in one step — no manual padding required:

```go
strings.Join(strings.Fields(s), " ")
```

## Approach

1. Bug: prepended a stray leading space: `return " " + collapsed`.
2. Fix: return `collapsed` directly — Fields already trims and collapses.
3. No leading/trailing whitespace remains.

## Solution

```go
import "strings"

func Normalize(s string) string {
	collapsed := strings.Join(strings.Fields(s), " ")
	return collapsed
}
```

## Walkthrough

"  hello   world  ": Fields -> ["hello","world"] -> Join -> "hello world".

## Pitfalls

- `Fields` already handles leading/trailing and multiple spaces; don't re-add
  any.
- `Fields("")` is empty, so `Join` yields `""` — correct.
- Tabs and newlines are whitespace too and collapse to a single space.

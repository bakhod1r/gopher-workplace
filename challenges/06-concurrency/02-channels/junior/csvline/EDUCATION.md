# Joining a Stream

## Intuition

Some operations are inherently non-streaming: joining needs to know whether
a field is the last one, which a stream cannot tell you until it ends. So
you drain into a slice and let `strings.Join` handle the boundaries.

## Approach

1. Collect the channel's fields into a `[]string`.
2. Return `strings.Join(parts, sep)`.

## Solution

```go
import "strings"

func JoinFields(fields <-chan string, sep string) string {
	parts := []string{}
	for field := range fields {
		parts = append(parts, field)
	}
	return strings.Join(parts, sep)
}
```

## Walkthrough

For `"a"`, `"b"` with `sep == ","`: the slice becomes `["a" "b"]` and
`strings.Join` produces `"a,b"`. One field yields just that field; an empty
record yields `""`.

## Pitfalls

- Appending `sep` after every field leaves a trailing delimiter on the line.
- `+=` in a loop is O(n²) for wide records; `strings.Join` (or `strings.Builder`) is linear.
- The empty record must return `""`, not `sep`.

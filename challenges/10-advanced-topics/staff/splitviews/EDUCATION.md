# Fields As Views, Into The Caller's Slice

## Intuition

Splitting moves no data — it only decides boundaries. The only allocation a splitter needs is the slice of headers, and even that can be the caller's if the API asks for it.

## Approach

1. Return `dst` unchanged for an empty line.
2. Walk `i` to `len(line)` inclusive, treating the end as a separator.
3. Append `line[start:i:i]` and move `start` past the separator.

## Solution

```go
// Fields appends each sep-separated field of line to dst as a view and
// returns the extended slice.
//
// The fields share line's storage, and dst lets the caller reuse the header
// slice, so a steady-state call allocates nothing at all.
//
// Examples:
//
// 	Fields(nil, []byte("a,b"), ',') => [][]byte{"a", "b"}
func Fields(dst [][]byte, line []byte, sep byte) [][]byte {
	if len(line) == 0 {
		return dst
	}
	start := 0
	for i := 0; i <= len(line); i++ {
		if i < len(line) && line[i] != sep {
			continue
		}
		dst = append(dst, line[start:i:i])
		start = i + 1
	}
	return dst
}
```

## Walkthrough

For "ab,cd", the boundaries are at 2 and 5. The views are `line[0:2:2]` and `line[3:5:5]`, so appending to the first must reallocate rather than overwrite the comma.

## Pitfalls

- `line[start:i]` without the third index — the field's capacity runs to the end of the line.
- Stopping at `len(line)-1`, which drops the last field.
- Treating an empty line as one empty field; the spec adds nothing.

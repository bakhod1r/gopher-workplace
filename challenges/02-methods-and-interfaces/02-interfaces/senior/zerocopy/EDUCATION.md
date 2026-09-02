# Zero Copy Slicing

## Intuition

A slice is a view, not a container. Returning `data[a:b]` costs nothing but ties the field's validity to the buffer's — which is a fine trade when the buffer outlives the parse, and a bug when it does not.

## Approach

1. Count the separators first so the result slice is allocated exactly once.
2. Walk the buffer, appending `data[start:i]` at each separator.
3. Append the final segment after the loop.
4. `CopyFields` parses, then copies each field into its own array.

## Solution

```go
func (p *ZeroCopyParser) Fields(data []byte) [][]byte {
	if len(data) == 0 {
		return nil
	}

	n := 1
	for _, b := range data {
		if b == p.Sep {
			n++
		}
	}

	out := make([][]byte, 0, n)
	start := 0
	for i := 0; i < len(data); i++ {
		if data[i] == p.Sep {
			out = append(out, data[start:i])
			start = i + 1
		}
	}
	out = append(out, data[start:])
	return out
}
```

## Walkthrough

`TestFieldsAliasInput` mutates the source buffer and expects the field to change — the aliasing is the documented behaviour, not an accident. `CopyFields` is the escape hatch for callers who need independence.

## Pitfalls

- `string(data[start:i])` per field — correct output, one allocation per field, and the whole optimisation is gone.
- Handing zero-copy fields to code that outlives the buffer (a pooled or reused array), which corrupts them later.
- Growing `out` from zero capacity, which reintroduces per-parse reallocation.

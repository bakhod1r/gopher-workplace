# Reader Interface

## Intuition

A stream interface must be able to say "nothing left". Returning `(chunk, ok)` lets the consumer loop without knowing the source's size.

## Approach

1. Return `"", false` when `pos` has reached the end.
2. Compute `end = pos + Chunk`, clamped to `len(Data)`.
3. Slice `Data[pos:end]`, set `pos = end`, return the chunk with `true`.
4. `ReadAll` loops until `ok` is false, concatenating chunks.

## Solution

```go
func (s *StringSource) Read() (string, bool) {
	if s.pos >= len(s.Data) {
		return "", false
	}
	end := s.pos + s.Chunk
	if end > len(s.Data) {
		end = len(s.Data)
	}
	chunk := s.Data[s.pos:end]
	s.pos = end
	return chunk, true
}

func ReadAll(s Source) string {
	out := ""
	for {
		chunk, ok := s.Read()
		if !ok {
			return out
		}
		out += chunk
	}
}
```

## Walkthrough

`"hello"` with `Chunk: 2` yields `"he"`, `"ll"`, then `end = 6` clamps to 5 giving `"o"`, then `pos == 5` ends the stream.

## Pitfalls

- Not clamping `end` — `s.Data[4:6]` panics with a slice bounds error.
- Returning `true` on the final empty read, which makes `ReadAll` loop forever.
- A value receiver: `pos` never advances and the loop never ends.

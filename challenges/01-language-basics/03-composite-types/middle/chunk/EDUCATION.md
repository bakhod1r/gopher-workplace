# Windowing a slice

## Intuition

Chunking steps an index by `size` and takes a sub-slice each time, clamping the
end so the final (short) chunk fits:

```go
for i := 0; i < len(xs); i += size {
	end := i + size
	if end > len(xs) { end = len(xs) }
	out = append(out, xs[i:end])
}
```

## Approach

1. Start with an empty non-nil [][]int.
2. If size<=0 return it.
3. Step i by size across xs.
4. Clamp end to len(xs); append xs[i:end].
5. Return result.

## Solution

```go
func Chunk(xs []int, size int) [][]int {
	out := [][]int{}
	if size <= 0 {
		return out
	}
	for i := 0; i < len(xs); i += size {
		end := i + size
		if end > len(xs) {
			end = len(xs)
		}
		out = append(out, xs[i:end])
	}
	return out
}
```

## Walkthrough

xs=[1,2,3,4,5] size=2: i=0 append [1,2]; i=2 append [3,4]; i=4 end clamps to 5 append [5]. Result [[1,2],[3,4],[5]].

## Pitfalls

- The chunks **share** `xs`'s backing array (sub-slices, not copies).
- Guard `size <= 0` to avoid an infinite loop.
- `min(i+size, len(xs))` (Go 1.21+) replaces the manual clamp.

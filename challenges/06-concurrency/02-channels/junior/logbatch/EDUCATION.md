# Collecting a Stream Into a Slice

## Intuition

Channels stream; slices hold. Batching means draining the stream into
memory. Because `range` stops only at close, this pattern needs a reader
that closes — otherwise the shipper never flushes.

## Approach

1. Initialise `batch` as an empty, non-nil slice.
2. `range` over `lines`, appending each one.
3. Return `batch`.

## Solution

```go
func CollectLines(lines <-chan string) []string {
	batch := []string{}
	for line := range lines {
		batch = append(batch, line)
	}
	return batch
}
```

## Walkthrough

For `"a"`, `"b"`, close: append `"a"`, append `"b"`, loop exits, return
`["a" "b"]`. For a container that exits without logging, the loop body never
runs and the initial `[]string{}` is returned.

## Pitfalls

- `var batch []string` returns `nil` for the silent case, which fails a `DeepEqual` against `[]string{}`.
- Ranging over a channel nobody closes hangs the shipper.
- Pre-allocating with `make([]string, n)` and then appending prepends `n` empty lines.

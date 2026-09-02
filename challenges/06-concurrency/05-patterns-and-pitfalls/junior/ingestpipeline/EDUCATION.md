# Composing a Pipeline

## Intuition

A pipeline is a chain of `<-chan` handoffs. Each link runs concurrently, but
because every record passes through one channel at a time, the order the
caller sees is exactly the order the file was read in.

## Approach

1. Feeder goroutine: send each line, `defer close(src)`.
2. Parse goroutine: `for line := range src { parsed <- parse(line) }`, `defer close(parsed)`.
3. Caller: `for rec := range parsed`, append when `isError(rec)`, return the slice.

## Solution

```go
func IngestPipeline(lines []string, parse func(string) string, isError func(string) bool) []string {
	src := make(chan string)
	go func() {
		defer close(src)
		for _, line := range lines {
			src <- line
		}
	}()

	parsed := make(chan string)
	go func() {
		defer close(parsed)
		for line := range src {
			parsed <- parse(line)
		}
	}()

	var out []string
	for rec := range parsed {
		if isError(rec) {
			out = append(out, rec)
		}
	}
	return out
}
```

## Walkthrough

For `err disk, info ok, err io` the parse stage emits `ERR DISK, INFO OK,
ERR IO`. The caller keeps the first and third. When the feeder closes `src`,
the parse stage's range ends and it closes `parsed`, ending the caller's
range — shutdown cascades in one direction.

## Pitfalls

- Appending to the result slice from inside a goroutine without synchronisation — that is a data race.
- Returning before draining `parsed`, which leaves the parse goroutine blocked on a send forever (a goroutine leak).
- Buffering the channels to "make it faster": it changes nothing here and hides a missing close.

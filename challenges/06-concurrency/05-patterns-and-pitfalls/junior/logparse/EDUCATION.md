# Pipeline Stages

## Intuition

A stage owns exactly one channel (its output) and reads from one it does not
own (its input). Because the input's close ends the loop, closing propagates
down the whole ingest chain without any coordination code.

## Approach

1. Create the output channel.
2. Range over `lines` inside a goroutine, sending `parse(line)`.
3. `defer close(out)` so the close happens exactly once, when the range ends.

## Solution

```go
func ParseStage(lines <-chan string, parse func(string) string) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for line := range lines {
			out <- parse(line)
		}
	}()
	return out
}
```

## Walkthrough

With lines `warn disk, info ok`: the goroutine receives the first line, sends
`WARN DISK`, then the second, sends `INFO OK`. The tailer closes `lines`,
`range` exits, the deferred `close(out)` runs, and the indexing stage stops.

## Pitfalls

- Using a counted loop instead of `range` — a tail never tells you how many lines are coming.
- Closing `lines` from inside the stage: a stage never closes a channel it did not create.
- Calling `parse` in the calling goroutine, which turns the stage into a blocking function instead of a pipeline step.

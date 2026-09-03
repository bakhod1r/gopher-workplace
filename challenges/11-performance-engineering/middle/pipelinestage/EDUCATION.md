# Stages, Channels, And Shutdown

## Intuition

Each stage is a loop: read until the input closes, transform, send. Closing the output on the way out is what lets the next stage finish too.

## Approach

1. `Stage` launches a goroutine that ranges its input and closes its output with `defer`.
2. `Run` builds a source channel, chains the stages over it, and drains the last one.

## Solution

```go
func Stage(in <-chan int, f func(int) int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			out <- f(v)
		}
	}()
	return out
}

func Run(values []int, stages []func(int) int) []int {
	src := make(chan int)
	go func() {
		defer close(src)
		for _, v := range values {
			src <- v
		}
	}()
	var ch <-chan int = src
	for _, f := range stages {
		ch = Stage(ch, f)
	}
	out := make([]int, 0, len(values))
	for v := range ch {
		out = append(out, v)
	}
	return out
}
```

## Walkthrough

`defer close(out)` is the whole shutdown story: when `in` closes, the `range` ends, the deferred close runs, and the next stage's `range` ends in turn — all the way down to the collecting loop in `Run`.

## Pitfalls

- Forgetting `close(out)`, so every downstream stage blocks forever and the goroutines leak.
- Sending on the source channel from `Run`'s own goroutine before draining, which deadlocks on an unbuffered channel.
- Adding a fan-out in the middle of a pipeline and losing the ordering the single channel was providing.

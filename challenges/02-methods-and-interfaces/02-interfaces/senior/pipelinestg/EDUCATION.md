# Pipeline Stage

## Intuition

The pipeline convention is strict: the goroutine that writes a channel is the one that closes it, and it lives exactly as long as its input. Break either rule and you get a leak or a send on a closed channel.

## Approach

1. Both stages are pure functions of `v`; `DropOddStage` returns `emit` false for odd values.
2. `RunStage` creates `out`, starts one goroutine, and returns immediately.
3. The goroutine `defer close(out)` and ranges over `in`.
4. Send only when `emit` is true.

## Solution

```go
func (DoubleStage) Process(v int) (int, bool) { return v * 2, true }

func (DropOddStage) Process(v int) (int, bool) { return v, v%2 == 0 }

func RunStage(in <-chan int, s Stage) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			if r, emit := s.Process(v); emit {
				out <- r
			}
		}
	}()
	return out
}
```

## Walkthrough

Chaining works because each stage's output is another stage's input: when the source closes, stage 1's range ends, its deferred close fires, and the closure cascades down the chain.

## Pitfalls

- Forgetting `close(out)` — downstream `range` blocks forever.
- Closing `in` inside the stage; a stage never closes a channel it did not create.
- Sending the value even when `emit` is false, which silently disables the filter.

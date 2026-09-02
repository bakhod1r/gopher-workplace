# Repeat Generators

## Intuition

Generators do not need an input collection — they only need a rule for what to
emit and a rule for when to stop. Here the rule is "the same probe URL, count
times", and the loop bound is the stop condition.

## Approach

1. Create the output channel.
2. In a goroutine, loop `count` times sending the probe URL.
3. `defer close(out)` so the consumer's `range` terminates.

## Solution

```go
func HealthPings(endpoint string, count int) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := 0; i < count; i++ {
			out <- endpoint + "/health"
		}
	}()
	return out
}
```

## Walkthrough

For `count = 2`: the goroutine sends the URL, the consumer receives it, the
goroutine sends it again, the consumer receives, the loop condition fails, and
`close(out)` ends the consumer's range with exactly two probes collected.

## Pitfalls

- Looping forever and relying on the consumer to stop reading — that leaks the goroutine.
- Skipping the close when `count <= 0`, which hangs the consumer on an empty round.
- Building the URL once outside the goroutine is fine; rebuilding it per send is only wasted work, not a bug.

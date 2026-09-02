# Signal Channels

## Intuition

When a goroutine only needs to say "I'm ready", the payload is irrelevant.
`chan struct{}` makes that explicit: the empty struct occupies zero bytes,
so the channel carries pure timing information. Receiving once per launched
shard is the simplest possible barrier.

## Approach

1. Clamp `n` to 0.
2. Make `ready := make(chan struct{})`.
3. Launch `n` goroutines, each sending `struct{}{}`.
4. Receive `n` times, counting.
5. Return the count.

## Solution

```go
func WaitForShards(n int) int {
	if n < 0 {
		n = 0
	}
	ready := make(chan struct{})
	for i := 0; i < n; i++ {
		go func() {
			ready <- struct{}{}
		}()
	}

	count := 0
	for i := 0; i < n; i++ {
		<-ready
		count++
	}
	return count
}
```

## Walkthrough

For `n = 3`: three goroutines each block on their send until the main loop
receives. Three receives later, `count` is 3 and every shard has warmed up.
For `n = 0` nothing is launched and nothing is received.

## Pitfalls

- Receiving fewer than `n` times leaks the goroutines still blocked on their sends.
- Receiving more than `n` times blocks forever — the service never starts.
- `chan bool` works but implies the value matters; `chan struct{}` states that it does not.

# Unbuffered Channels as a Rendezvous

## Intuition

An unbuffered channel has no storage. A send parks until some goroutine
receives, and vice versa. That makes it a *synchronisation point*, not a
queue: after `probes <- i` returns, you know the upstream has the value.
Alternating two such channels gives a strict turn-taking protocol.

## Approach

1. Clamp `rounds` to 0.
2. Make two unbuffered channels, `probes` and `acks`.
3. Start a goroutine that loops `rounds` times: receive from `probes`, send on `acks`.
4. In the sidecar loop, send on `probes`, trace `"probe"`, receive from `acks`, trace `"ack"`.
5. Return the trace.

## Solution

```go
func Probe(rounds int) []string {
	if rounds < 0 {
		rounds = 0
	}
	probes := make(chan int)
	acks := make(chan int)

	go func() {
		for i := 0; i < rounds; i++ {
			<-probes
			acks <- i
		}
	}()

	trace := []string{}
	for i := 0; i < rounds; i++ {
		probes <- i
		trace = append(trace, "probe")
		<-acks
		trace = append(trace, "ack")
	}
	return trace
}
```

## Walkthrough

For `rounds = 1`: the sidecar sends on `probes` and blocks; the upstream
receives, so both continue. The sidecar traces `"probe"` and blocks on
`<-acks`; the upstream sends, the sidecar receives and traces `"ack"`.

## Pitfalls

- Using a single channel for both directions lets one goroutine receive its own send.
- Buffering the channels breaks the lock-step and the trace can interleave wrongly.
- A mismatched loop count leaves one side blocked forever — a deadlock.

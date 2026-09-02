# Ordering Is the Producer's Job

## Intuition

A channel is strictly FIFO: it never reorders. Any ordering the renderer
sees was chosen by the sender. This is worth internalising, because it
also means that with *several* senders no order is guaranteed at all — the
interleaving is up to the scheduler.

## Approach

1. Make an unbuffered `chan string`.
2. In a goroutine, loop `i` from `len(lines)-1` down to `0`, sending `lines[i]`.
3. `close(ch)` after the loop.
4. `range` in the renderer, appending into a non-nil slice.
5. Return it.

## Solution

```go
func TailReverse(lines []string) []string {
	ch := make(chan string)
	go func() {
		for i := len(lines) - 1; i >= 0; i-- {
			ch <- lines[i]
		}
		close(ch)
	}()

	out := []string{}
	for line := range ch {
		out = append(out, line)
	}
	return out
}
```

## Walkthrough

For `["a","b","c"]`: the goroutine sends `"c"`, `"b"`, `"a"`, each handed
over as the renderer receives it, then closes. The collected view is
`["c" "b" "a"]`.

## Pitfalls

- `i > 0` as the loop condition drops the oldest line; it must be `i >= 0`.
- Sending without a goroutine on an unbuffered channel deadlocks immediately.
- Two producer goroutines would destroy the ordering guarantee.

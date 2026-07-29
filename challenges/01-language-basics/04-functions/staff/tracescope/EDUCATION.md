# Scoping defer to a loop iteration

## Intuition

Deferred calls run at function return, so a `defer` in a loop body accumulates to the end; wrap each iteration in its own function literal to make the defer fire per iteration.

## Approach

1. `defer` fires at **function** exit, not loop-iteration exit.
2. Wrap each iteration's body in an immediately-invoked func so its defer runs per iteration.

## Solution

```go
import "fmt"

func Trace(n int) (log []string) {
	for i := 0; i < n; i++ {
		func(k int) {
			log = append(log, fmt.Sprintf("start%d", k))
			defer func() { log = append(log, fmt.Sprintf("end%d", k)) }()
		}(i)
	}
	return
}
```

## Walkthrough

The plain `defer` inside the loop queues all `end` logs to the function's end, giving `[start0 start1 end1 end0]`. An IIFE per iteration scopes the defer so `end` follows each `start`.

## Pitfalls

- A loop-body `defer` runs at function exit, not each iteration's end.
- `func(){ ...; defer end() }()` per iteration gives interleaved order.

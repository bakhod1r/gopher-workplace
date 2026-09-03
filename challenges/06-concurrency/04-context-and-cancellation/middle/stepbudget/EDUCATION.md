# Sequential Step Budget

## Intuition

Two contexts are in play per iteration: the request's, which decides whether to continue at all, and the step's, which bounds that one call. Keeping them distinct is what lets a step's stragglers be cleaned up without killing the whole chain.

## Approach

1. Track how many steps ran.
2. At the top of each iteration, return early on `ctx.Err()`.
3. Derive `stepCtx` with `WithCancel`, call the step, then `cancel()`.
4. Count the step, and return its error if there is one.
5. After the loop, return `ctx.Err()` (nil for a live context).

## Solution

```go
func RunSteps(ctx context.Context, steps []Step) (int, error) {
	ran := 0
	for _, step := range steps {
		if err := ctx.Err(); err != nil {
			return ran, err
		}

		stepCtx, cancel := context.WithCancel(ctx)
		err := step(stepCtx)
		cancel()

		ran++
		if err != nil {
			return ran, err
		}
	}
	return ran, ctx.Err()
}
```

## Walkthrough

With three steps and a live parent: step one gets its own child, runs, and its context is cancelled before step two is derived — so no goroutine it left behind survives. If the parent is cancelled during step one, the next iteration's `ctx.Err()` returns `context.Canceled`, the chain stops at one step run, and steps two and three never touch the network.

## Pitfalls

- `defer cancel()` inside the loop: every step's context stays alive until the whole chain ends.
- Passing the parent straight to the step, so a hung step's goroutines outlive it.
- Checking `ctx.Err()` only at the start, which keeps calling upstream after the client hung up.

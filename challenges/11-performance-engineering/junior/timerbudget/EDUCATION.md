# The Budget Is Per Request, Not Per Call

## Intuition

Latency budgets are additive. Every component's share is its own cost multiplied by how often the request goes through it.

## Approach

1. `Cost` guards and multiplies.
2. `Headroom` subtracts from the budget without clamping.
3. `Fits` checks the headroom's sign.

## Solution

```go
func Cost(nsPerOp int64, callsPerRequest int) int64 {
	if nsPerOp <= 0 || callsPerRequest <= 0 {
		return 0
	}
	return nsPerOp * int64(callsPerRequest)
}

func Headroom(nsPerOp int64, callsPerRequest int, budgetNS int64) int64 {
	return budgetNS - Cost(nsPerOp, callsPerRequest)
}

func Fits(nsPerOp int64, callsPerRequest int, budgetNS int64) bool {
	return Headroom(nsPerOp, callsPerRequest, budgetNS) >= 0
}
```

## Walkthrough

Building `Fits` on `Headroom` and `Headroom` on `Cost` means the guards exist in one place, and the inclusive boundary falls out of `>= 0`.

## Pitfalls

- Clamping `Headroom` at zero, which hides a 10x overrun behind "no headroom".
- A strict `> 0` in `Fits`, rejecting a component that lands exactly on budget.
- Comparing per-call latency against a per-request budget and concluding everything is fine.

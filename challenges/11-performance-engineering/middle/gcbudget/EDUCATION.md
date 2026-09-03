# Trading Memory For GC Time

## Intuition

Collections happen when the heap grows by `GOGC` percent. Allow twice the growth and you run half as many collections over the same allocation volume — at the price of a heap that is bigger by exactly that ratio.

## Approach

1. `CPUFraction` divides GC nanoseconds by the window's total CPU nanoseconds.
2. `TuneGOGC` refuses non-positive inputs and satisfied budgets, then scales and caps.
3. `MemoryCost` applies the ratio.

## Solution

```go
func CPUFraction(gcNS, windowNS int64, cores int) float64 {
	if windowNS <= 0 || cores <= 0 {
		return 0
	}
	return float64(gcNS) / (float64(windowNS) * float64(cores))
}

func TuneGOGC(currentGOGC int, currentFraction, targetFraction float64, maxGOGC int) (int, bool) {
	if currentGOGC <= 0 || targetFraction <= 0 {
		return 0, false
	}
	if currentFraction <= targetFraction {
		return 0, false
	}
	next := int(math.Round(float64(currentGOGC) * (currentFraction / targetFraction)))
	next = max(next, currentGOGC)
	if maxGOGC > 0 {
		next = min(next, maxGOGC)
	}
	return next, true
}

func MemoryCost(gogc int) (float64, bool) {
	if gogc <= 0 {
		return 0, false
	}
	return 1 + float64(gogc)/100, true
}
```

## Walkthrough

`math.Round` rather than a plain conversion matters here: `0.30/0.10` is 2.9999999999999996 in binary floating point, and truncating it would suggest a GOGC of 299 instead of 300.

Refusing to "tune" a budget that is already met is the interesting guard: the formula would happily return a *lower* GOGC, which trades away memory headroom to buy CPU nobody asked for.

## Pitfalls

- Forgetting the core count, which overstates the GC's share by up to the parallelism of the machine.
- Raising GOGC without checking the memory ceiling, turning a CPU problem into an out-of-memory kill.
- Treating the inverse relationship as exact; it is a planning estimate, and the real answer comes from measuring again.

# Trading Memory For GC Time

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

The Go collector has one real knob. `GOGC` sets how much the heap may grow before the next cycle, and it buys CPU with RAM: doubling it roughly halves the collector's CPU share and roughly doubles the heap. Knowing the exchange rate turns "the GC is eating 20% of our CPU" into a decision instead of a complaint.

## Task

Implement the three functions in [gcbudget.go](gcbudget.go):

1. `CPUFraction` returns GC CPU time over the window's total CPU capacity — `windowNS * cores` — with a non-positive window or core count giving 0.
2. `TuneGOGC` suggests `currentGOGC * (currentFraction / targetFraction)`, rounded to the nearest whole value and capped at `maxGOGC`, and reports `false` when the budget is already met or the inputs are non-positive.
3. `MemoryCost` returns the heap multiplier `1 + GOGC/100`, reporting `false` for a non-positive GOGC.

## Examples

**Example 1:**

```
Input:  CPUFraction(100ms, 1s, 8)
Output: 0.0125
```

**Example 2:**

```
Input:  TuneGOGC(100, 0.20, 0.10, 1000)
Output: 200, true
```

**Example 3:**

```
Input:  MemoryCost(200)
Output: 3, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **GC cost is inverse to headroom** | More room between collections means fewer collections over the same work. |
| 2 | **Cores are part of the denominator** | 100ms of GC in a second is 10% of one core and 1.25% of eight. |
| 3 | **Every tuning is a trade** | The CPU you save comes straight out of the memory budget — say so when you suggest it. |

## Topics used again

Float ratios, floors and ceilings, multiple return values, guards.

## Hint

`TuneGOGC` should refuse when `currentFraction <= targetFraction` — the "improvement" would be a regression in memory for nothing.

## Validate

```bash
make verify
```

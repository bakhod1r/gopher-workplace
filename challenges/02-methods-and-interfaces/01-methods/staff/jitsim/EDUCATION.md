# JIT Tier-Up

## Intuition

Compiling costs time, so a JIT only pays it for code that will run again. The
cheapest possible predictor of "will run again" is "has run a lot already" — a
counter and a threshold. Everything else in a tiered runtime (on-stack
replacement, deoptimisation, inlining decisions) is built on top of this
bookkeeping.

## Approach

1. Bump the counter for this op.
2. If it reached the threshold, set the sticky compiled flag.
3. Report the tier that actually served this execution.

## Solution

```go
func (j *JIT) Execute(op string) Tier {
	j.counts[op]++
	if j.counts[op] >= j.Threshold {
		j.compiled[op] = true
	}
	if j.compiled[op] {
		return Compiled
	}
	return Interpreted
}
```

## Walkthrough

With `Threshold: 3` and one op:

| execution | count after | compiled | returned |
|-----------|-------------|----------|----------|
| 1 | 1 | no | interp |
| 2 | 2 | no | interp |
| 3 | 3 | yes | jit |
| 4 | 4 | yes | jit |

The second test proves the counters are keyed: `"b"`'s single execution does not
help `"a"` reach two, and `"a"`'s promotion leaves `"b"` cold.

## Pitfalls

- **Checking before incrementing.** Promotion slips by one execution, and
  `New(1)` never fires at all.
- **`==` instead of `>=`.** Works while every op passes through the threshold
  exactly, but a counter that is bumped elsewhere (or a threshold lowered at
  runtime) skips the promotion permanently.
- **Deriving the tier from the count instead of the flag.** Equivalent here, but
  the flag is what makes the decision sticky and independently inspectable —
  which is what `IsCompiled` reports.
- **Nil maps.** Constructing `JIT{}` directly panics on the first write; `New`
  exists for that reason.

## What real tiering adds

Go's own compiler is ahead-of-time, but PGO plays a similar role: profile
counts decide what gets inlined. Runtimes with real JITs (V8, HotSpot) add
several tiers, on-stack replacement so a long-running loop can switch tier
mid-execution, and deoptimisation when a speculative assumption breaks.

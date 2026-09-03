# Dropping The Wrong End

## Intuition

`s[:len(s)-n]` and `s[n:]` are the same length and opposite meanings: one throws away the tail, the other throws away the head. Warm-up lives at the head.

## Approach

1. Slice from the warmup count to the end.

## Solution

```go
rest := samples[warmup:]
```

## Walkthrough

For `[500, 5, 5, 5, 5, 5]` with one warmup sample, the bug keeps `[500, 5, 5, 5, 5]` and reports a mean of 104 — the cold first run survives and one perfectly good steady-state sample is thrown away instead.

## Pitfalls

- Any length-only assertion, which passes for both versions.
- Trimming the tail on purpose to drop a shutdown effect, then forgetting which end you meant.
- Discarding so many samples that the remaining set is too small to be stable.

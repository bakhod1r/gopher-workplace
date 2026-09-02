# Peak Latency

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The API gateway's latency monitor consumes per-request durations for one
rolling window and reports the worst one. A window with no requests has
nothing to report.

## Task

Implement `PeakLatency` in [latencypeak.go](latencypeak.go) so that:

1. It drains `samples` until the window closes.
2. It returns the maximum latency received and `true`.
3. It returns `0, false` for an empty window — and must work for clock-skewed negative samples.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PeakLatency(30, 90, 40)
Output: 90, true
```

**Example 2:**

```
Input:  PeakLatency() // closed, empty
Output: 0, false
```

**Example 3:**

```
Input:  PeakLatency(-5, -2)
Output: -2, true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Streaming reduction** | One pass, constant memory, no slice needed. |
| 2 | **Seeding the accumulator** | Seed from the first sample, not from `0`. |
| 3 | **Empty-window signal** | The `bool` result reports "no requests". |

## Hint

Do not start `peak` at `0` — an all-negative window would then report `0`.
Seed it from the first sample using the `seen` flag.

## Validate

```bash
make verify
```

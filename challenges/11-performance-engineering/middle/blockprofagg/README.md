# Reading A Sampled Block Profile

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

The block profile records where goroutines wait: channel sends, mutex acquisitions, network reads. It is off by default and sampled when on — `runtime.SetBlockProfileRate(n)` records roughly one event per `n` nanoseconds of blocking. Forget the scaling factor and every number you read is off by a factor of a million.

## Task

Implement the three functions in [blockprofagg.go](blockprofagg.go):

1. `Scale` multiplies a sampled wait by the profile rate; a rate at or below 1 scales by 1.
2. `Totals` aggregates scaled waits per site, dropping non-positive waits.
3. `FractionBlocked` divides the scaled total by a wall-clock window; a non-positive window gives 0, and the result may exceed 1.

## Examples

**Example 1:**

```
Input:  Scale(100, 1000000)
Output: 100000000
```

**Example 2:**

```
Input:  Totals([{a 100} {b 20} {a 50}], 1)
Output: {a:150 b:20}
```

**Example 3:**

```
Input:  FractionBlocked([{a 100} {b 100}], 1, 100)
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sampling needs scaling** | The recorded events stand for the ones that were not recorded. |
| 2 | **Blocked time is not wall time** | Ten goroutines blocked for a second is ten seconds of blocking in one second. |
| 3 | **The profile is off by default** | Enabling it costs throughput, which is why nobody has it on until they need it. |

## Topics used again

Map aggregation, int64 scaling, float division, guards.

## Hint

`Totals` and `FractionBlocked` should both go through `Scale` rather than repeating the multiplication.

## Validate

```bash
make verify
```

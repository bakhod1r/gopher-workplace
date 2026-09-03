# Nanoseconds Into Cycles

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

"40 ns/op" means nothing until you convert it into work the CPU actually does. At 3 GHz that is 120 cycles — enough for a cache miss and change, far too much for an integer add. Converting first tells you whether a number is suspicious before you open a profiler.

## Task

Implement both functions in [cyclebudget.go](cyclebudget.go):

1. `CyclesPerOp` converts ns/op to cycles at `ghz` gigahertz; a non-positive input gives `0`.
2. `Verdict` classifies a cycle count: `"register"` below 10, `"l1"` below 100, `"memory"` below 1000, `"syscall"` at 1000 or above.
3. A non-positive cycle count is `"register"`.

## Examples

**Example 1:**

```
Input:  CyclesPerOp(10, 3)
Output: 30
```

**Example 2:**

```
Input:  Verdict(30)
Output: "l1"
```

**Example 3:**

```
Input:  Verdict(1000)
Output: "syscall"
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A cycle is the unit of work** | Nanoseconds are clock-speed dependent; cycles are comparable across machines. |
| 2 | **Order-of-magnitude intuition** | Register ops are single digits, an L1 hit a few, main memory hundreds. |
| 3 | **Half-open boundaries** | Each band includes its lower bound and excludes the next one. |

## Topics used again

Float arithmetic, `switch` with conditions, guards.

## Hint

`ns * ghz` is the whole conversion; a `switch` with no tag reads better than nested ifs.

## Validate

```bash
make verify
```

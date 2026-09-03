# Keeping The Optimizer Honest

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A benchmark that computes a value and throws it away is a benchmark the compiler may delete outright — you measure an empty loop and celebrate. The fix is a package-level sink the optimizer cannot prove is unused.

## Task

Implement both functions in [benchsink.go](benchsink.go):

1. `SumTo(n)` returns `0+1+...+(n-1)`; a non-positive `n` returns `0`.
2. `Consume(v)` stores `v` in the package-level `Sink` and returns the value it replaced.

## Examples

**Example 1:**

```
Input:  SumTo(4)
Output: 6
```

**Example 2:**

```
Input:  SumTo(-3)
Output: 0
```

**Example 3:**

```
Input:  Sink is 0, then Consume(7)
Output: 0 returned, Sink is now 7
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Dead-code elimination** | A result no one observes can be compiled away, loop and all. |
| 2 | **Package-level sink** | An exported variable is observable outside the package, so the store must happen. |
| 3 | **Cheap sink** | Storing an `int` costs almost nothing, so it does not distort the measurement. |

## Topics used again

Package-level variables, loops.

## Hint

`Consume` needs the old value before it overwrites it.

## Validate

```bash
make verify
```

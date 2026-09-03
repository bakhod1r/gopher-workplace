# Pad To The Line, Not To The Word

**Level:** staff
**Topic:** 10-advanced-topics / 04-unsafe-package

## Context

A sharded counter scales negatively: eight workers are slower than one. The counters are separate variables, no lock is held, and the race detector is silent.

## Task

Implement [paddedcounters.go](paddedcounters.go):

1. Give each worker its own `Slot` and have it increment `N` `iters` times.
2. Join, then return the sum of the slots.
3. Non-positive `workers` or negative `iters` return 0.
4. The slot stride must be a full cache line — the padding follows `unsafe.Sizeof`, not a magic number.

Replace the stub body in [paddedcounters.go](paddedcounters.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Run(4, 1000)
Output: 4000
```

**Example 2:**

```
Input:  sizeof(Slot)
Output: 64
```

_Explanation:_ One counter per line.

**Example 3:**

```
Input:  &s[1] - &s[0]
Output: 64
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **False sharing** | Coherence works per line, so neighbours on one line serialise. |
| 2 | **unsafe.Sizeof in a constant expression** | The padding adapts if the counter's type changes. |
| 3 | **Array stride** | The element size is the struct's size, which is what separates the slots. |
| 4 | **Join before summing** | `wg.Wait()` is what makes the counters safe to read. |

## Hint

The padded type is written for you. The body is allocate, fan out, join, fold.

## Validate

```bash
make verify
```

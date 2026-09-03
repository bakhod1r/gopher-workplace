# The Same Work Every Run

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

A benchmark fed fresh random data measures a different workload on every run, and the noise swamps the change you are trying to evaluate. Seeded, reproducible input makes two runs comparable.

## Task

Implement `FixedInput` in [benchfixedinput.go](benchfixedinput.go):

1. Start from `state = seed` and step it with `state = state*1664525 + 1013904223`.
2. Emit the new state as each of the `n` values, in order.
3. A non-positive `n` returns an empty, non-nil slice.

## Examples

**Example 1:**

```
Input:  FixedInput(1, 1)
Output: [1664526]
```

**Example 2:**

```
Input:  FixedInput(42, 8) twice
Output: identical slices
```

**Example 3:**

```
Input:  FixedInput(1, 0)
Output: [] (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reproducible input** | Comparing two runs is only meaningful if they did the same work. |
| 2 | **Unsigned wraparound** | `uint32` arithmetic wraps by definition, which is what makes the generator legal and portable. |
| 3 | **Build input outside the timed loop** | Generation cost is setup, not the operation under test. |

## Topics used again

`make` with a capacity hint, integer overflow rules.

## Hint

One accumulator, one loop, no `math/rand`.

## Validate

```bash
make verify
```

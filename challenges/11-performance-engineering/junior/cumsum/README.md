# The Cumulative Column

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Flat time says which function burns the CPU; cumulative time says which function is *responsible* for it. `main` has almost no flat time and 100% cum. The rule: every frame on a sample's stack is credited with that sample's value — but a recursive function must not be credited once per frame.

## Task

Implement `CumSum` in [cumsum.go](cumsum.go):

1. Credit each function on a sample's stack with the sample's `Value`.
2. Credit a function that appears several times in the same stack only once for that sample.
3. Ignore samples with a non-positive `Value` or an empty stack; return an empty, non-nil map when nothing survives.

## Examples

**Example 1:**

```
Input:  [{[main a b] 5}]
Output: {main:5 a:5 b:5}
```

**Example 2:**

```
Input:  [{[main a] 3} {[main b] 4}]
Output: {main:7 a:3 b:4}
```

**Example 3:**

```
Input:  [{[main rec rec rec] 6}]
Output: {main:6 rec:6}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cum is flat plus callees** | The whole stack is charged, which is why cum totals exceed the wall time. |
| 2 | **Recursion double-counting** | Crediting per frame inflates recursive functions without bound. |
| 3 | **A set per sample** | Deduplicating within one sample, not across samples, is the whole trick. |

## Topics used again

Maps as sets, nested loops, `range`.

## Hint

Build a `map[string]bool` of the frames already credited — one per sample, reset for the next.

## Validate

```bash
make verify
```

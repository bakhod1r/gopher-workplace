# Only The Frame On The CPU

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Given the same sample list, cumulative time charges every frame on the stack while self time charges exactly one: the leaf, the function that was actually running when the profiler interrupted. Two aggregations, one input — and reading a profile means knowing which one you are looking at.

## Task

Implement `SelfTime` in [funcselftime.go](funcselftime.go):

1. Credit only the last frame of each stack with the sample's `Value`.
2. Accumulate across samples.
3. Ignore samples with a non-positive `Value` or an empty stack; return an empty, non-nil map when nothing survives.

## Examples

**Example 1:**

```
Input:  [{[main a b] 5}]
Output: {b:5}
```

**Example 2:**

```
Input:  [{[main a] 3} {[main b a] 4} {[main b] 1}]
Output: {a:7 b:1}
```

**Example 3:**

```
Input:  [{[main] 2}]
Output: {main:2}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Leaf is the running frame** | Callers are waiting, not working, at the instant of the sample. |
| 2 | **Self totals equal the profile total** | Unlike cum, self time partitions the samples exactly once. |
| 3 | **The same function can be a leaf anywhere** | Its self time accumulates across every stack it terminates. |

## Topics used again

Slice indexing, maps, `range`.

## Hint

`s.Stack[len(s.Stack)-1]` — guard the empty stack before you index it.

## Validate

```bash
make verify
```

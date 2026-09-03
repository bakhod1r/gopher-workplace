# A Flame Graph With Impossible Call Paths

## Intuition

Sorting looks like canonicalisation, and for a set it would be. A call stack is not a set: `a;b` and `b;a` describe opposite relationships.

## Approach

1. Use the stack as it is, without reordering it.

## Solution

```go
frames := s.Stack
```

## Walkthrough

With the bug, `[a b]` and `[b a]` both fold to `a;b`, so their values are summed into one line — the flame graph shows a single path carrying the weight of two, and half of the arrows in it point the wrong way.

## Pitfalls

- Sorting or deduplicating a stack anywhere it is used as an identity.
- Joining with a separator that can appear inside a symbol name, which merges paths the same way.
- Trusting an aggregation because the totals still add up; they do, into the wrong buckets.

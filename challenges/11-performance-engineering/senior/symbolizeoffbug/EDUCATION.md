# The Function Entry Point Blamed On Its Neighbour

## Intuition

The search finds a boundary, and the code then steps back one symbol. If the boundary already includes the matching symbol, stepping back lands on the previous function.

## Approach

1. Search for the first symbol whose start is strictly greater than the address.

## Solution

```go
i := sort.Search(len(table), func(i int) bool { return table[i].Start > addr })
```

## Walkthrough

For `addr = 200` and a symbol starting exactly at 200, the buggy predicate is already true at that index, so `i-1` points at `a` — the previous function absorbs every sample that landed on `b`'s entry point.

## Pitfalls

- Off-by-one predicates in any predecessor search; the fix is one character and the symptom is subtle.
- Testing only with addresses in the middle of a range, which both versions get right.
- Blaming the linker or the profiler when the attribution looks slightly shifted.

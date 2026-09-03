# Every Profile Says `main` Is Hot

## Intuition

The stack is stored caller-first, so index 0 is the outermost frame — the one that has been waiting the longest. The frame on the CPU is at the other end.

## Approach

1. Credit the last frame instead of the first.

## Solution

```go
frame := s.Stack[len(s.Stack)-1]
```

## Walkthrough

With the bug, every sample under `main` credits `main`, so the flat column becomes a second cumulative column and the actual hot leaf never appears at all.

## Pitfalls

- Assuming the stack orientation instead of checking it; both conventions are common.
- Indexing before the length guard, which panics on an empty stack.
- Reading a flat profile whose top entry is the entry point without suspecting the aggregation.

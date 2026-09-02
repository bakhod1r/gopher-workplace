# Ring Buffer

**Level:** junior  
**Topic:** 03-generics

## Context

A crash reporter keeps only the last N log lines in memory, whatever type those lines are.

## Task

Implement the stub(s) in [ringgen.go](ringgen.go):

1. Implement `NewRing`, `Add`, and `Items`.
2. `Add` drops the oldest element once the ring is full.
3. `Items` returns a copy, oldest first; a zero-size ring stores nothing.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewRing[int](2); Add(1); Add(2); Add(3); Items()
Output: []int{2, 3}
```

**Example 2:**

```
Input:  NewRing[int](2); Items()
Output: []int{}
```

**Example 3:**

```
Input:  NewRing[int](0); Add(1); Items()
Output: []int{}
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generic types** | `type Stack[T any] struct { ... }` parameterises the type itself, not just a function. |
| 2 | **Bounded growth** | Trimming on every `Add` keeps memory flat regardless of input volume. |
| 3 | **Defensive copies** | Reused from earlier: returning the internal slice would let callers corrupt the ring. |

## Hint

Append first, then trim to size — one branch, no modular arithmetic.

## Validate

```bash
make verify
```

# Escape Audit

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A hot function slowed down after an innocent refactor. The values now escape to the heap and the GC is doing the extra work.

## Task

Implement the stub(s) in [escapeaudit.go](escapeaudit.go):

1. Implement `Sum` on `*StackAgg` so the aggregator's state does not escape per call.
2. Implement `SumValues`, which sums a slice without boxing each element into an interface.
3. Implement `SumBoxed`, the deliberately boxing version, for comparison.
4. Constraint: `SumValues` must allocate zero times per call; `SumBoxed` is expected to allocate — the tests assert both.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SumValues([1 2 3])
Output: 6, zero allocations
```

**Example 2:**

```
Input:  SumBoxed([]any{1, 2, 3})
Output: 6, allocations from boxing
```

**Example 3:**

```
Input:  an empty slice
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Escape analysis** | A value escapes when the compiler cannot prove its lifetime ends with the frame. |
| 2 | **Interface boxing** | Storing a non-pointer in an interface generally allocates. |
| 3 | **Measured cost claims** | Reused: `AllocsPerRun` grades the claim. |

## Hint

Passing a concrete `[]int` keeps everything on the stack; converting each element to `any` heap-allocates.

## Validate

```bash
make verify
```

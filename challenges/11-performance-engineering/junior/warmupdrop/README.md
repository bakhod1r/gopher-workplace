# Throwing Away The First Rounds

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

The first iteration of anything is unrepresentative: cold caches, lazily built tables, a page fault per new heap page. `go test -bench` handles this by running the loop repeatedly and reporting the steady state; when you collect samples yourself you have to do the same thing by hand.

## Task

Implement both functions in [warmupdrop.go](warmupdrop.go):

1. `Drop` discards the first `n` samples without modifying or aliasing the input.
2. A non-positive `n` keeps everything; an `n` at or beyond the length leaves an empty, non-nil slice.
3. `StableMean` drops the warm-up and averages the rest; nothing left gives `0`.

## Examples

**Example 1:**

```
Input:  Drop([9 1 1], 1)
Output: [1 1]
```

**Example 2:**

```
Input:  Drop([1 2], 5)
Output: [] (non-nil)
```

**Example 3:**

```
Input:  StableMean([100 2 4], 1)
Output: 3
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Warm-up is a one-off cost** | Averaging it in makes a fast steady state look slow. |
| 2 | **Reslicing aliases** | `samples[n:]` shares the array, so writes to the result reach the caller. |
| 3 | **Dropping everything is legal** | It is a small sample set, not an error. |

## Topics used again

`slices.Clone`, slice bounds, `min`, float division.

## Hint

Clamp `n` into `[0, len(samples)]` before you slice with it.

## Validate

```bash
make verify
```

# Five Stages, Five Copies

**Level:** staff  
**Topic:** 03-generics

## Context

A five-stage transform over million-row batches allocates five million-element buffers per batch. The output is correct; the allocation rate is five times what the design budgeted and GC pauses show it.

## Task

Fix the single planted bug in [pipelinerebuildbug.go](pipelinerebuildbug.go):

1. Find and fix the single bug so the pipeline reuses one buffer across stages.
2. The input slice must not be modified and the results must not change.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Pipeline([1,2], inc, double)
Output: [4 6]
```

**Example 2:**

```
Input:  input after the call
Output: unchanged
```

**Example 3:**

```
Input:  allocations for 5 stages
Output: one buffer, not five
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Reuse the buffer** | Element-wise stages can be applied in place once the input has been copied out. |
| 2 | **Copy once, at the boundary** | One defensive copy protects the caller; the rest are waste. |
| 3 | **Backing-array aliasing** | A returned slice that shares storage lets the caller mutate the source. |

## Hint

Count the calls to `make` for a five-stage run.

## Validate

```bash
make verify
```

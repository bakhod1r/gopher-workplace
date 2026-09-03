# Stages, Channels, And Shutdown

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

A pipeline turns a sequential transform into an overlapping one: while stage two works on value one, stage one is already on value two. The throughput win is real and so is the classic failure — a stage that never learns its input is finished, leaving a goroutine parked on a channel read for the life of the process.

## Task

Implement both functions in [pipelinestage.go](pipelinestage.go):

1. `Stage` runs `f` over everything arriving on `in`, sending results onward and closing its output when `in` closes.
2. `Run` chains the stages in order and collects the results, preserving input order.
3. No goroutine may still be running when `Run` returns; no stages means the identity transform, and no values gives an empty non-nil slice.

## Examples

**Example 1:**

```
Input:  Run([1 2 3], [double inc])
Output: [3 5 7]
```

**Example 2:**

```
Input:  Run([1 2], nil)
Output: [1 2]
```

**Example 3:**

```
Input:  a closed empty input channel
Output: Stage's output channel closes immediately
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Closing propagates** | `range` over a channel ends when it closes, which is how shutdown travels down the pipeline. |
| 2 | **Channels preserve order** | One channel between two stages is FIFO, so no sorting is needed. |
| 3 | **A leaked stage leaks forever** | A goroutine blocked on a channel nobody will close is never collected. |

## Topics used again

Channels, directional channel types, `range` over a channel, goroutines, `close`.

## Hint

`Run` can feed the source channel from its own goroutine, close it, and then drain the final stage on the main one.

## Validate

```bash
make verify
```

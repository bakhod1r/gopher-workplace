# Merge Channels Without Leaving Goroutines Behind

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A fan-in helper merges worker channels. The consumer often stops early on the first error, and the process accumulates blocked forwarder goroutines until it is restarted.

## Task

Implement [fanin.go](fanin.go):

1. Return a channel carrying every value from `ins`.
2. Close it once every input is drained.
3. Every goroutine must exit when `done` is closed, even if the consumer has stopped reading.
4. Zero inputs closes the output immediately.

Replace the stub body in [fanin.go](fanin.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Merge(done, send(1,2), send(3))
Output: a channel yielding 1, 2 and 3
```

**Example 2:**

```
Input:  inputs drained
Output: out is closed
```

**Example 3:**

```
Input:  consumer abandons out, done closed
Output: no goroutine left behind
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A blocked send is a live goroutine** | It holds its stack and everything its frame references. |
| 2 | **select with a cancellation channel** | Makes every send abandonable. |
| 3 | **WaitGroup then close** | The closer must run after all forwarders, in its own goroutine. |
| 4 | **Only the sender closes** | The forwarders send; the extra goroutine closes. |

## Hint

Each forwarder's send is the thing that can block forever. What else should it be able to do?

## Validate

```bash
make verify
```

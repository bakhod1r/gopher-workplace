# Fan-In With Shutdown

**Level:** staff
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A merge of several producer channels leaked a goroutine per input whenever the consumer stopped reading early.

## Task

Implement the stub(s) in [fanin.go](fanin.go):

1. Implement `Merge`, which fans several input channels into one output.
2. Every forwarding goroutine must exit when its input closes **or** when `done` is closed, and the output must close exactly once.
3. Constraint: `-race` clean, no goroutine leak after an early consumer exit — the test asserts the goroutine count returns to baseline.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  merge two channels, drain fully
Output: every value, output closed
```

**Example 2:**

```
Input:  close done after one read
Output: all forwarders exit
```

**Example 3:**

```
Input:  merge nothing
Output: an immediately closed output
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fan-in** | One goroutine per input, one WaitGroup, one closer. |
| 2 | **done-channel shutdown** | The consumer's escape hatch, checked in every send. |
| 3 | **Close ownership** | Reused: only the goroutine that created the channel closes it. |

## Hint

Send with a `select` on `out <- v` and `<-done`, and close `out` from a goroutine that waits on the WaitGroup.

## Validate

```bash
make verify
```

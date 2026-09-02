# Dead Letter Drain

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

When a consumer group is retired, whatever is left in its dead-letter queue
has to be drained and counted for the shutdown report. The producer is a
goroutine that only exits once its last send is received, so a partial drain
would leave it — and the shutdown — stuck.

## Task

Implement `DrainDeadLetters` in [deadletters.go](deadletters.go) so that:

1. It ranges over `msgs` until the channel is closed.
2. It counts every message received, discarding the payload.
3. It returns the count, which is zero for an already-empty queue.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  DrainDeadLetters(channel of 3 messages)
Output: 3
```

**Example 2:**

```
Input:  DrainDeadLetters(channel of 1 message)
Output: 1
```

**Example 3:**

```
Input:  DrainDeadLetters(closed empty channel)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Draining** | `for range ch {}` receives until close and unblocks the producer. |
| 2 | **Range over channel** | The loop ends exactly when the channel is closed. |
| 3 | **Producer liveness** | An undrained producer stays blocked on its next send forever. |

## Hint

`for range msgs` without a value variable is legal Go and is the idiomatic
drain — increment a counter in the body.

## Validate

```bash
make verify
```

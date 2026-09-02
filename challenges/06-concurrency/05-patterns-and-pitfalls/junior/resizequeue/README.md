# Resize Queue

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The image service runs a fixed pool of resize workers behind the upload queue.
The pool size is a deployment knob: it caps how many resizes run at once no
matter how long the queue grows, which is what keeps the box from thrashing
during a traffic spike.

## Task

Implement `ResizeQueue` in [resizequeue.go](resizequeue.go) so that:

1. It creates a jobs channel and a results channel, and starts exactly `workers` goroutines.
2. Each worker ranges over jobs and sends `resize(key)` to results.
3. The caller queues all uploads, closes the jobs channel, waits for the workers, closes results, and returns the drained keys sorted ascending (`nil` for an empty queue).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ResizeQueue([]string{"a", "b"}, 2, resize)
Output: []string{"a-512", "b-512"}
```

**Example 2:**

```
Input:  ResizeQueue([]string{"z"}, 4, resize)
Output: []string{"z-512"}
```

**Example 3:**

```
Input:  ResizeQueue(nil, 3, resize)
Output: nil
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Worker pool** | Jobs channel + results channel + fixed worker count. |
| 2 | **Close ordering** | close(jobs) → wg.Wait() → close(results); any other order breaks. |
| 3 | **Backpressure** | An unbuffered jobs channel means the producer waits for a free worker. |

## Hint

Write the shutdown sequence first — `close(jobs)`, `wg.Wait()`,
`close(results)` — then fill in the workers around it.

## Validate

```bash
make verify
```

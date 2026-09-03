# Bound The Work In Flight

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A crawler starts a goroutine per URL. With a large frontier it opens ten thousand sockets, exhausts the file descriptor limit, and takes the rest of the process with it.

## Task

Implement [semaphore.go](semaphore.go):

1. Take one slot, waiting until one is free.
2. Report false and take nothing when `done` is closed first.
3. Never allow more than the configured number of holders.

Replace the stub body in [semaphore.go](semaphore.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  NewSem(2), two Acquires
Output: true, true
```

**Example 2:**

```
Input:  a third Acquire while full
Output: blocks, then false when done closes
```

**Example 3:**

```
Input:  Release while someone waits
Output: the waiter proceeds
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **A buffered channel is a semaphore** | Its capacity is the permit count. |
| 2 | **Send to acquire, receive to release** | The buffer's occupancy is the held count. |
| 3 | **Cancellable waiting** | `select` over the acquire and `done`. |
| 4 | **Acquire nothing on cancellation** | Returning false must not leave a slot taken. |

## Hint

Two cases in one `select`: taking a slot, and giving up.

## Validate

```bash
make verify
```

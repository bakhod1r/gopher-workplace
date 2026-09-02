# Round-Robin Router

**Level:** staff
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A router spreads incoming messages across a fixed pool of workers. Round-robin
is the simplest fair policy: keep a cursor, deliver to the worker it points at,
then advance it and wrap.

## Task

Implement `Route` on `*Router` in [actorrouter.go](actorrouter.go):

1. Send `msg` to `r.workers[r.idx].Inbox`.
2. Advance `r.idx` to the next worker, wrapping with `% len(r.workers)`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  2 workers; Route(1)
Output: worker 0 receives 1
```

**Example 2:**

```
Input:  Route(2)
Output: worker 1 receives 2
```

**Example 3:**

```
Input:  Route(3)
Output: worker 0 receives 3   (cursor wrapped)
```

_Explanation:_ the modulo is what makes the third message land back on worker 0.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Cursor state on the receiver** | `idx` must persist between calls, so the receiver is a pointer. |
| 2 | **Modulo wrap** | `r.idx = (r.idx + 1) % len(r.workers)` keeps the index in range forever. |
| 3 | **Buffered inboxes** | Capacity 10 means `Route` never blocks in these tests. |

## Hint

Send first, then advance. Advancing first delivers message 1 to worker 1 and
shifts the whole sequence.

## Validate

```bash
make verify
```

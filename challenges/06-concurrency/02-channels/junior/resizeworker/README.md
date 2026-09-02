# Scale Request

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The image resize service asks its worker for the scaled width of a retina
asset: the handler sends the source width and blocks until the worker sends
the scaled width back on a reply channel.

## Task

Implement `ScaleRequest` in [resizeworker.go](resizeworker.go) so that:

1. It creates two unbuffered channels: one for the request, one for the reply.
2. A goroutine receives the source width, doubles it, and sends the result on the reply channel.
3. The handler sends `width`, then returns the value it receives back.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ScaleRequest(640)
Output: 1280
```

**Example 2:**

```
Input:  ScaleRequest(0)
Output: 0
```

**Example 3:**

```
Input:  ScaleRequest(-3)
Output: -6
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Request/reply** | Two channels give a synchronous round trip. |
| 2 | **Unbuffered handoff** | Each send completes only when the other side receives. |
| 3 | **No close needed** | One value each way — nothing to signal termination. |

## Hint

Start the worker goroutine **before** sending on `req`, otherwise the
unbuffered send has no receiver and deadlocks.

## Validate

```bash
make verify
```

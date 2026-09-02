# Try Record

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The metrics client must never slow down the request path. If the sample
buffer is full the sample is dropped and the caller is told, rather than the
handler stalling on a full channel.

## Task

Implement `TryRecord` in [metricdrop.go](metricdrop.go) so that:

1. It uses `select` with a send case and a `default`.
2. It returns `true` when the sample was recorded.
3. It returns `false` immediately when the send would block — it must never wait.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  TryRecord(cap-1 empty buffer, 5)
Output: true
```

**Example 2:**

```
Input:  TryRecord(cap-1 full buffer, 5)
Output: false
```

**Example 3:**

```
Input:  TryRecord(unbuffered buffer, 5)
Output: false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Non-blocking send** | `select` with `default` around a send case. |
| 2 | **Send readiness** | A buffered send is ready while `len < cap`; an unbuffered one needs a waiting receiver. |
| 3 | **Lossy by design** | Dropping a sample is better than blocking a request. |

## Hint

Same shape as a non-blocking receive, but the case is `buffer <- sample`
instead of `<-buffer`.

## Validate

```bash
make verify
```

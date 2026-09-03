# Merge Order Feeds

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

Every region publishes its own order feed. The checkout reconciler wants one
channel to range over, not a slice of channels. Fan-in gives it that — but the
merged channel can only be closed once, and only after the *last* region has
finished. Whoever closes too early truncates the reconciliation.

## Task

Implement `MergeOrderFeeds` in [ordermerge.go](ordermerge.go) so that:

1. It returns a channel immediately, without draining the feeds first.
2. One goroutine per feed forwards every id onto the merged channel.
3. A single closer goroutine waits for all forwarders and then closes the merged channel exactly once.
4. With zero feeds the returned channel is closed straight away, so ranging over it terminates.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MergeOrderFeeds(feed[eu-1, eu-2], feed[us-1])
Output: channel yielding eu-1, eu-2, us-1 in some order, then closed
```

**Example 2:**

```
Input:  MergeOrderFeeds(feed[], feed[ap-1])
Output: channel yielding ap-1, then closed
```

**Example 3:**

```
Input:  MergeOrderFeeds()
Output: an already-closed, empty channel
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Close ownership** | The merged channel has many senders, so no sender may close it. |
| 2 | **WaitGroup + closer goroutine** | `wg.Wait()` then `close` runs off the caller's goroutine so the function can return early. |
| 3 | **Directional types** | Feeds arrive as `<-chan string`; the result is handed back receive-only. |
| 4 | **Per-goroutine loop variable** | Each forwarder must capture its own feed. |

## Hint

Return the channel before any forwarding has happened. The only goroutine
allowed to call `close(merged)` is the one that first calls `wg.Wait()`.

## Validate

```bash
make verify
```

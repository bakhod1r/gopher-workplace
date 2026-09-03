# Fan Out a Chat Room

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

A chat room has one incoming message feed and N connected clients. Each client
runs its own goroutine that ranges over a personal channel, so every client
must receive every message, in publication order, and every client channel must
eventually close — otherwise the client goroutine leaks when the room shuts down.

## Task

Implement `FanOutRoom` in [chatfanout.go](chatfanout.go) so that:

1. It creates one unbuffered channel per subscriber and returns them as `[]<-chan string`, immediately.
2. A single broadcaster goroutine ranges the room feed and sends each message to every subscriber, in order.
3. When the room feed closes, the broadcaster closes every subscriber channel.
4. With `subscribers == 0` it returns an empty slice and still drains the feed, so the publisher is never left blocked.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  FanOutRoom(feed[hi, bye], 2)
Output: 2 channels, each yielding hi then bye, then closed
```

**Example 2:**

```
Input:  FanOutRoom(feed[hi], 1)
Output: 1 channel yielding hi, then closed
```

**Example 3:**

```
Input:  FanOutRoom(feed[hi, bye], 0)
Output: [] — and the feed is drained anyway
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fan-out** | One producer, N consumers, each getting a full copy of the stream. |
| 2 | **Close as broadcast end** | Closing every output is what lets each client's `range` return. |
| 3 | **Directional conversion** | A `chan string` you own converts to `<-chan string` for the caller. |
| 4 | **Backpressure** | Unbuffered outputs mean the slowest subscriber paces the room. |

## Hint

Keep two slices: the bidirectional `chan string` you send and close on, and the
`<-chan string` view you hand back. Return before the broadcaster has sent
anything — the caller has to be receiving for the unbuffered sends to complete.

## Validate

```bash
make verify
```

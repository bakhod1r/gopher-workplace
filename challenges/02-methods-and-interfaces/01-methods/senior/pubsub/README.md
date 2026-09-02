# Pub/Sub

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

Where an event bus calls listeners directly, a pub/sub hands each subscriber a
channel. Publishing means sending on every channel registered for the topic.
The subscriber map is shared, so reads must be guarded by the read lock.

## Task

Implement `Publish` on `*PubSub` in [pubsub.go](pubsub.go):

1. Take `ps.mu.RLock()` (and release it).
2. Iterate `ps.subs[topic]` and send `msg` on each channel.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  two subscribers on "news"; Publish("news", "hello")
Output: both channels receive "hello"
```

**Example 2:**

```
Input:  Publish to a topic with no subscribers
Output: no-op
```

**Example 3:**

```
Input:  Publish twice
Output: each subscriber receives both messages, in order (buffer size 10)
```

_Explanation:_ the channels are buffered, so publishing does not block.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`sync.RWMutex`** | Many concurrent `Publish` calls may hold `RLock` at once; `Subscribe`'s `Lock` excludes them all. |
| 2 | **`defer` for unlocking** | `defer ps.mu.RUnlock()` releases even if a send panics. |
| 3 | **Directional channels** | `Subscribe` returns `<-chan string`, so callers can only receive. |

## Hint

Buffered channels (capacity 10) mean a plain `ch <- msg` will not deadlock in
the tests. Do **not** take `Lock` — that would serialize publishers needlessly.

## Validate

```bash
make verify
```

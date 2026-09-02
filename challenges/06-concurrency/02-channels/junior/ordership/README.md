# Ship Order IDs

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The checkout service publishes order ids to a Kafka producer goroutine over
a channel. The publisher is the sender, so the publisher is what closes the
channel once the batch of orders has been handed over.

## Task

Implement `ShipOrderIDs` in [ordership.go](ordership.go) so that:

1. It sends each element of `ids` on `out`, preserving order.
2. It closes `out` exactly once, after the last send.
3. A `nil` or empty batch results in an immediately closed channel.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  ShipOrderIDs(out, []int{101, 102})
Output: out: 101, 102 then closed
```

**Example 2:**

```
Input:  ShipOrderIDs(out, nil)
Output: out: closed, no values
```

**Example 3:**

```
Input:  ShipOrderIDs(out, []int{7})
Output: out: 7 then closed
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Send-only parameter** | `chan<- int` allows sending and closing, never receiving. |
| 2 | **Sender closes** | Only the sender may close, and only once. |
| 3 | **`range` over a slice** | Feeds the channel in batch order. |

## Hint

Loop over the batch sending each id, then `close(out)` **after** the loop —
not inside it.

## Validate

```bash
make verify
```

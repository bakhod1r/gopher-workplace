# Cancellable Price Feed

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

A price feed has no natural end: it quotes until the subscriber goes away. If
the producing goroutine only knows how to send, every closed subscription
leaves one goroutine parked forever on `out <- price`. Over a trading day that
is thousands of leaked goroutines.

## Task

Implement `PriceFeed` in [pricefeed.go](pricefeed.go) so that:

1. It returns a channel and streams `base`, `base+1`, `base+2`, ... on it.
2. Every send is wrapped in a `select` that also watches `done`.
3. When `done` is closed, the goroutine returns and closes the output channel.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PriceFeed(done, 100), take 3
Output: 100, 101, 102
```

**Example 2:**

```
Input:  PriceFeed(done, 0), take 5
Output: 0, 1, 2, 3, 4
```

**Example 3:**

```
Input:  close(done)
Output: the feed stops and its channel closes
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Goroutine lifetime** | A producer must never outlive the consumer that gave it a reason to run. |
| 2 | **Cancellable send** | `select` with `done` makes a send abandonable. |
| 3 | **Infinite generator** | The loop has no bound; cancellation is the only exit. |

## Hint

The infinite loop needs an exit that does not depend on the consumer reading:
that is the `case <-done: return` next to the send.

## Validate

```bash
make verify
```

# Transactional Mail Semaphore

**Level:** middle
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The notification service sends transactional email through a provider that rejects anything above a fixed number of concurrent connections. Every queued message must still be attempted, so the sender uses a buffered channel as a counting semaphore instead of shrinking the goroutine count.

## Task

Implement the stubbed function in [mailsemaphore.go](mailsemaphore.go) so that:

1. Send every message, writing each provider response into its input position.
2. Never run more than `limit` sends at the same time.
3. A `limit` of zero or less means unlimited.
4. Return an empty slice for an empty input.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SendAll([]string{"m1","m2"}, 1, send)
Output: ["sent:m1" "sent:m2"]
```

**Example 2:**

```
Input:  SendAll([]string{"m1"}, 5, send)
Output: ["sent:m1"]
```

**Example 3:**

```
Input:  SendAll(nil, 2, send)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | Counting semaphore | A buffered channel of capacity `n`: a send takes a slot, a receive returns one, and the buffer size is the limit. |
| 2 | Acquire before spawning | Taking the slot in the loop — not inside the goroutine — caps the number of live goroutines too. |
| 3 | Release with `defer` | `defer func(){ <-sem }()` returns the slot on every exit path, panic included. |
| 4 | Disjoint index writes | Each goroutine owns `out[i]`, so no mutex is needed and order is preserved for free. |

## Hint

`sem := make(chan struct{}, limit)`. Send into it before `go`, and receive from it in a `defer` inside the goroutine.

## Validate

```bash
make verify
```

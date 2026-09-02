# Email Send Semaphore

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The SMTP relay accepts a limited number of simultaneous connections, and
exceeding it gets the sender throttled. A *counting semaphore* built from a
buffered channel caps how many deliveries are in flight, while the campaign
still fans out across goroutines.

## Task

Implement `SendCampaign` in [emailqueue.go](emailqueue.go) so that:

1. It creates a buffered channel of capacity `limit` to act as the semaphore.
2. One goroutine per recipient acquires a slot (`sem <- struct{}{}`), calls `send(addr)`, and releases the slot.
3. The running total of the per-send results is guarded by a mutex and returned after `wg.Wait()`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  SendCampaign([]string{"a", "bb"}, 2, len)
Output: 3
```

**Example 2:**

```
Input:  SendCampaign([]string{"abcd"}, 1, len)
Output: 4
```

**Example 3:**

```
Input:  SendCampaign(nil, 3, len)
Output: 0
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Buffered channel as semaphore** | Capacity = permits; a send acquires, a receive releases. |
| 2 | **Mutex for shared state** | `total += n` from many goroutines needs a lock. |
| 3 | **Acquire/release symmetry** | `defer func() { <-sem }()` guarantees the permit comes back even on panic. |

## Hint

`sem <- struct{}{}` blocks once `limit` permits are taken. Release with a
deferred receive so an early return can never strand a permit.

## Validate

```bash
make verify
```

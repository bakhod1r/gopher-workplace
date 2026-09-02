# Async Fetch

**Level:** senior
**Topic:** 02-methods-and-interfaces → 01-methods

## Context

A service needs to start work now and let the caller decide when to wait for it.
The idiom is a method that launches a goroutine and hands back a
`<-chan struct{}`: the caller receives from it, and the closed channel is the
"done" signal.

## Task

Implement `FetchAsync` on `*Fetcher` in [asyncfetch.go](asyncfetch.go):

1. Create a `chan struct{}`.
2. Launch a goroutine that calls `f.Fetch(id)`.
3. Close the channel when `Fetch` finishes.
4. Return the channel immediately.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  f.FetchAsync("abc"); <-done
Output: f.Result == "data: abc"
```

**Example 2:**

```
Input:  the returned channel, before Fetch finishes
Output: a receive on it blocks
```

**Example 3:**

```
Input:  a second receive after the channel is closed
Output: returns immediately (a closed channel never blocks)
```

_Explanation:_ `close` is a broadcast — every current and future receive is released.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Goroutines in methods** | A method may launch a goroutine exactly like a function; the receiver is captured by the closure. |
| 2 | **Completion channel** | `close(done)` signals all receivers; a send would release only one. |
| 3 | **Directional return type** | `<-chan struct{}` means the caller can only receive, never close. |

## Hint

Return the channel *before* the work is done — that is the point. Close it
inside the goroutine, after `Fetch` returns, with `defer close(done)`.

## Validate

```bash
make verify
```

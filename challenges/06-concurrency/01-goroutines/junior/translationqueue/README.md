# Translation Queue

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A translation service is billed per character, not per byte. Before dispatching
a queue of messages to the vendor, the client counts the characters in each one
to build the cost estimate. Messages are counted concurrently and the estimate
keeps queue order.

## Task

Implement `CharCounts` in [translationqueue.go](translationqueue.go) so that:

1. Return a slice of counts the same length as `messages`.
2. Count `i` is the number of runes in `messages[i]`; a multi-byte character counts as one.
3. Count each message in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CharCounts([]string{"go", "añb"})
Output: [2 3]
```

**Example 2:**

```
Input:  CharCounts([]string{"日本"})
Output: [2]
```

**Example 3:**

```
Input:  CharCounts(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Runes, not bytes** | `utf8.RuneCountInString` bills characters; `len` would overcharge every non-ASCII message. |

## Hint

`utf8.RuneCountInString` counts without allocating a `[]rune`, and it is a pure
function, so concurrent calls are safe.

## Validate

```bash
make verify
```

# Stream Pages

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The export endpoint pages through a large result set. A producer goroutine
emits page numbers lazily so the exporter never holds more than one page
number ahead of the writer.

## Task

Implement `StreamPages` in [pagefeed.go](pagefeed.go) so that:

1. A goroutine sends `1 .. n` on an **unbuffered** channel and closes it.
2. The exporter collects the page numbers with `range`.
3. `n <= 0` yields an empty, non-nil slice and the goroutine still closes the channel.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  StreamPages(3)
Output: [1 2 3]
```

**Example 2:**

```
Input:  StreamPages(1)
Output: [1]
```

**Example 3:**

```
Input:  StreamPages(0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Goroutine producer** | Concurrency lets an unbuffered channel work without deadlock. |
| 2 | **`close` inside the goroutine** | The sender closes, from the goroutine that sends. |
| 3 | **`range` consumer** | The exporter blocks and collects until close. |

## Hint

Because the producer runs concurrently, `make(chan int)` with no buffer is
enough — each send waits for the exporter's receive.

## Validate

```bash
make verify
```

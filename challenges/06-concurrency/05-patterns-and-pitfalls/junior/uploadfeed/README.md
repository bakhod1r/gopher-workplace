# Upload Feed

**Level:** junior
**Topic:** 06-concurrency → 05-patterns-and-pitfalls

## Context

The media service receives a batch of object keys whenever a user finishes an
upload, and a downstream thumbnail pipeline consumes them one at a time. The
first link in that pipeline is a *generator*: it turns the batch slice into a
stream on a channel so the workers can pull keys as they free up.

## Task

Implement `UploadFeed` in [uploadfeed.go](uploadfeed.go) so that:

1. It creates a channel and returns its receive-only end immediately.
2. A goroutine sends each key of the batch in order.
3. The channel is closed once the batch is exhausted, so a `range` over it terminates.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  UploadFeed([]string{"a.jpg", "b.jpg"})
Output: "a.jpg", "b.jpg" then closed
```

**Example 2:**

```
Input:  UploadFeed([]string{"only.png"})
Output: "only.png" then closed
```

**Example 3:**

```
Input:  UploadFeed(nil)
Output: closed immediately, no keys
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Generator pattern** | A function that returns a channel it alone writes to and closes. |
| 2 | **Channel ownership** | The goroutine that creates a channel is the one that closes it. |
| 3 | **Directional channels** | Returning `<-chan string` stops consumers from sending or closing. |

## Hint

Create the channel, start a goroutine with `defer close(out)`, and return
the channel *before* the goroutine finishes — never after.

## Validate

```bash
make verify
```

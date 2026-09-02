# Preview Orders

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The admin UI shows a preview of the first few orders on a long live feed.
It must stop at the preview limit, and also stop cleanly if the feed ends
before the limit is reached.

## Task

Implement `PreviewOrders` in [previewfeed.go](previewfeed.go) so that:

1. It receives at most `limit` order ids from `feed`.
2. It stops early — without blocking — if the feed closes before `limit` ids arrive.
3. `limit <= 0` returns an empty, non-nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  PreviewOrders(1,2,3 | limit 2)
Output: [1 2]
```

**Example 2:**

```
Input:  PreviewOrders(1 | limit 5)
Output: [1]
```

**Example 3:**

```
Input:  PreviewOrders(1,2 | limit 0)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Bounded receive** | A counted loop instead of `range`. |
| 2 | **Comma-ok in a loop** | `ok == false` is the signal to `break`. |
| 3 | **Two stop conditions** | Whichever comes first: the limit or the close. |

## Hint

`range` cannot stop at a count. Use `for i := 0; i < limit; i++` with a
comma-ok receive and `break` when `ok` is false.

## Validate

```bash
make verify
```

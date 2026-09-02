# Slugifier

**Level:** junior
**Topic:** 06-concurrency → 01-goroutines

## Context

A CMS publishes a batch of drafts and needs a URL slug for each one: the title
lowercased, with runs of whitespace collapsed into single hyphens. Titles are
independent, so the batch is slugified concurrently.

## Task

Implement `Slugs` in [slugifier.go](slugifier.go) so that:

1. Return a slice of slugs the same length as `titles`.
2. Slug `i` is `titles[i]` split on whitespace, joined with `"-"`, and lowercased.
3. A title that is entirely whitespace produces the empty string.
4. Slugify each title in its own goroutine, joined with a `sync.WaitGroup`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Slugs([]string{"Hello World"})
Output: [hello-world]
```

**Example 2:**

```
Input:  Slugs([]string{"   "})
Output: []
```

**Example 3:**

```
Input:  Slugs(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`go` statement** | `go f(x)` starts a goroutine; the caller keeps running and does not wait. |
| 2 | **`sync.WaitGroup`** | `wg.Add(1)` before each launch, `defer wg.Done()` inside, `wg.Wait()` in the parent. |
| 3 | **Loop-variable capture** | Pass the index and the element in as goroutine parameters instead of reading the loop variable. |
| 4 | **Pure string functions** | `strings.Fields`, `Join` and `ToLower` only read their arguments, so concurrent calls need no locking. |

## Hint

`strings.Fields` already collapses runs of whitespace and never returns empty
fields, so `Join` plus `ToLower` finishes the job.

## Validate

```bash
make verify
```

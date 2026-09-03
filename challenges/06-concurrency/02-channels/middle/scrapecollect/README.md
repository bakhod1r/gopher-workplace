# Collect Scrapes on a Budget

**Level:** middle
**Topic:** 06-concurrency → 02-channels

## Context

The metrics collector fans out a scrape to every target and reads the replies
off one channel. A wedged target must not wedge the scrape interval, so the
collector gets a budget: whatever arrived by then is written as a partial
scrape, and the run is marked incomplete.

## Task

Implement `CollectScrapes` in [scrapecollect.go](scrapecollect.go) so that:

1. It receives from `scrapes` until it holds `want` samples, then returns them with `true`.
2. A `select` arm on a single deadline channel ends collection when the budget expires, returning what arrived so far with `false`.
3. If the scrape pool closes the channel first, it returns what arrived so far with `false`.
4. With `want <= 0` it returns an empty, non-nil slice and `true` without receiving anything.

The returned slice is always non-nil.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CollectScrapes(chan with 3 samples, want=3, budget=5s)
Output: [{api 12} {db 7} {web 3}], true
```

**Example 2:**

```
Input:  CollectScrapes(chan with 1 sample then closed, want=3, budget=5s)
Output: [{api 12}], false
```

**Example 3:**

```
Input:  CollectScrapes(silent chan, want=2, budget=20ms)
Output: [], false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Timeout arm** | `case <-deadline` in a `select` bounds a blocking receive. |
| 2 | **One deadline, not one per iteration** | `time.After` inside the loop restarts the budget on every sample. |
| 3 | **Comma-ok on receive** | `s, ok := <-scrapes` separates "a sample" from "the pool is done". |
| 4 | **Partial results** | A short read is data plus a flag, not an error that throws the data away. |

## Hint

Create the deadline channel **once**, before the loop, and select on the same
variable every iteration.

## Validate

```bash
make verify
```

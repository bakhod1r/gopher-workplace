# Merge Order Feeds

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The reconciliation report concatenates two order feeds: every order id from
the primary region first, then every id from the standby region. The order
between the two feeds is part of the report format.

## Task

Implement `MergeFeeds` in [orderfeeds.go](orderfeeds.go) so that:

1. It drains `primary` fully, appending each id.
2. It then drains `standby` fully, appending each id.
3. The report is primary ids followed by standby ids; two empty feeds give an empty, non-nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  MergeFeeds(primary: 1,2 | standby: 3)
Output: [1 2 3]
```

**Example 2:**

```
Input:  MergeFeeds(primary: empty | standby: 9)
Output: [9]
```

**Example 3:**

```
Input:  MergeFeeds(primary: 4 | standby: empty)
Output: [4]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Sequential drain** | Two `range` loops give a deterministic concatenation. |
| 2 | **Close per channel** | Each `range` needs its own feed to be closed. |
| 3 | **Receive-only params** | `<-chan int` for both feeds. |

## Hint

Two `range` loops, one after the other. No `select` needed — the report
format fixes the order.

## Validate

```bash
make verify
```

# Collect Log Lines

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

A log shipper reads lines from a container's stdout stream and buffers them
into one batch before making a single write to the log backend. The reader
closes the channel when the container exits.

## Task

Implement `CollectLines` in [logbatch.go](logbatch.go) so that:

1. It receives every line until `lines` is closed.
2. It returns them in arrival order.
3. A container that logged nothing yields an empty, non-nil slice — never `nil`.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  CollectLines("a", "b")
Output: ["a" "b"]
```

**Example 2:**

```
Input:  CollectLines() // closed, empty
Output: []
```

**Example 3:**

```
Input:  CollectLines("x")
Output: ["x"]
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`range` over a channel** | Terminates when the reader closes. |
| 2 | **`append` growth** | Building a batch incrementally from a stream. |
| 3 | **Non-nil empty slice** | `[]string{}` differs from `nil` under `reflect.DeepEqual`. |

## Hint

Start with `batch := []string{}` (not `var batch []string`) so a silent
container still produces a non-nil batch.

## Validate

```bash
make verify
```

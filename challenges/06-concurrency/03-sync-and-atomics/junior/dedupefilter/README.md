# Webhook Dedupe Filter

**Level:** junior  
**Topic:** 06-concurrency → 03-sync-and-atomics

## Context

A webhook receiver gets at-least-once delivery: the same event ID can arrive twice, on two connections, at the same instant. Each event must be processed exactly once, so the check-and-claim of an event ID has to be a single atomic step.

## Task

Implement the stubbed functions in [dedupefilter.go](dedupefilter.go) so that:

1. `Accept` returns true the first time an event ID is seen and false for redeliveries.
2. `Seen` reports whether an event ID has ever been accepted.
3. `Len` returns how many distinct events were accepted.

Do **not** change the function signatures or the tests.

## Examples

**Example 1:**

```
Input:  var d DedupeFilter; d.Accept("evt-1")
Output: true
```

**Example 2:**

```
Input:  d.Accept("evt-1"); d.Accept("evt-1")
Output: true, then false
```

**Example 3:**

```
Input:  d.Accept("evt-1"); d.Accept("evt-2"); d.Len()
Output: 2
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **sync.Map** | A concurrent map built for write-once / read-many keys, no external lock needed. |
| 2 | **LoadOrStore** | Stores only when absent and tells you which happened, in one atomic step. |
| 3 | **Range** | `Range(f)` walks the entries; return false from `f` to stop early. |

## Hint

`_, loaded := d.seen.LoadOrStore(id, struct{}{})` - `Accept` should return `!loaded`.

## Validate

```bash
make verify
go test -race ./...
```

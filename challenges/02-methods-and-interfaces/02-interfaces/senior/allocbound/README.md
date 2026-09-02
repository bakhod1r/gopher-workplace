# Allocation-Bounded Sink

**Level:** senior
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A hot ingest path appends millions of records. The current sink reallocates on every write and the GC never gets a break.

## Task

Implement the stub(s) in [allocbound.go](allocbound.go):

1. Implement `Write` on `*Sink` so it appends to a preallocated buffer.
2. Implement `NewSink`, which reserves capacity up front.
3. Implement `Fill`, which writes n records through the `Recorder` interface.
4. Constraint: filling a sink created with the right capacity must average **zero** allocations per write (checked with `testing.AllocsPerRun`).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  NewSink(1000); Fill(s, 1000)
Output: 1000 records, no reallocation
```

**Example 2:**

```
Input:  s.Len() after 3 writes
Output: 3
```

**Example 3:**

```
Input:  cap growth during Fill within capacity
Output: none
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Preallocation** | `make([]T, 0, n)` reserves the backing array once. |
| 2 | **Allocation accounting** | `testing.AllocsPerRun` turns a cost claim into a graded assertion. |
| 3 | **Interface dispatch cost** | Reused: the interface call itself must not box the argument. |

## Hint

`append` only allocates when it outgrows `cap`. Reserve once in `NewSink`.

## Validate

```bash
make verify
```

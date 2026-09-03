# Compare Bytes Without Making A String

**Level:** middle
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A protocol dispatcher checks the frame type with `strings.HasPrefix(string(frame), ...)`. Each check copies the whole frame to compare its first four bytes.

## Task

Implement [bytescompare.go](bytescompare.go):

1. Report whether `b` starts with `prefix`.
2. An empty prefix always matches; a prefix longer than `b` never does.
3. Zero allocations, whatever the size of `b`.

Replace the stub body in [bytescompare.go](bytescompare.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  HasPrefix([]byte("hello"), "he")
Output: true
```

**Example 2:**

```
Input:  HasPrefix([]byte("hello"), "hello!")
Output: false
```

_Explanation:_ A longer prefix cannot match.

**Example 3:**

```
Input:  HasPrefix([]byte("hello"), "")
Output: true
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **string([]byte) copies** | The conversion allocates because strings are immutable and slices are not. |
| 2 | **Indexing a string** | `prefix[i]` is a byte — no conversion needed to compare. |
| 3 | **Length guards** | Check the length before indexing, not after. |

## Hint

`b[i]` and `prefix[i]` are both bytes already.

## Validate

```bash
make verify
```

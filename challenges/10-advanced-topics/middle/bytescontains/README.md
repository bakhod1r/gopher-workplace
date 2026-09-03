# Search Bytes Without Building A String

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A filter checks every incoming frame for a marker with `strings.Contains(string(frame), marker)`. Each check copies the whole frame to look for eight bytes.

## Task

Implement [bytescontains.go](bytescontains.go):

1. Report whether `needle` appears in `haystack`.
2. An empty needle is present; a needle longer than the haystack is not.
3. Convert nothing — zero allocations.

Replace the stub body in [bytescontains.go](bytescontains.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Contains([]byte("hello"), "ell")
Output: true
```

**Example 2:**

```
Input:  Contains([]byte("hello"), "")
Output: true
```

_Explanation:_ The empty string is everywhere.

**Example 3:**

```
Input:  Contains([]byte("aaa"), "aab")
Output: false
```

_Explanation:_ A partial match must not stop the search.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **string([]byte) copies** | The conversion allocates because strings are immutable. |
| 2 | **Indexing both sides** | `haystack[i]` and `needle[j]` are both bytes already. |
| 3 | **Restarting after a partial match** | The outer loop advances by one, not past the failed match. |

## Hint

A first-byte check before the inner loop skips most positions cheaply.

## Validate

```bash
make verify
```

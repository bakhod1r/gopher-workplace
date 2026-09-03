# Split Without Copying The Text

**Level:** middle
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A CSV reader turns every field into a fresh string. For a wide file that is one small allocation per cell, and the cells outnumber everything else in the profile.

## Task

Implement [fieldsview.go](fieldsview.go):

1. Split `s` on `sep` and return every piece, including empty ones.
2. The pieces must be substrings of `s` — exactly one allocation per call, for the result slice.
3. An empty input yields one empty piece.

Replace the stub body in [fieldsview.go](fieldsview.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Fields("a,b,c", ',')
Output: ["a" "b" "c"]
```

**Example 2:**

```
Input:  Fields("a,,b", ',')
Output: ["a" "" "b"]
```

_Explanation:_ Empty fields are preserved.

**Example 3:**

```
Input:  Fields("", ',')
Output: [""]
```

_Explanation:_ No separator means one piece.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Substrings share bytes** | `s[a:b]` is a new header over the same immutable bytes. |
| 2 | **Counting before allocating** | `strings.Count` gives the exact piece count for one `make`. |
| 3 | **Immutability is what makes it safe** | Nobody can write through a string, so sharing needs no copy. |

## Hint

A string cannot be modified, so why would a piece of one need its own bytes?

## Validate

```bash
make verify
```

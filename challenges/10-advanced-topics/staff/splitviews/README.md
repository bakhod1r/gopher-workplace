# Fields As Views, Into The Caller's Slice

**Level:** staff
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A record splitter allocates a fresh `[][]byte` and copies every field on every line. At a million lines a second it is the whole cost of the parser.

## Task

Implement [splitviews.go](splitviews.go):

1. Append each `sep`-separated field of `line` to `dst` as a view, and return the extended slice.
2. Empty fields count, including leading and trailing ones; an empty line adds nothing.
3. Each view's capacity must stop at its own end, so appending to one cannot reach the next.
4. With room in `dst`, allocate nothing.

Replace the stub body in [splitviews.go](splitviews.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Fields(nil, []byte("a,bb,c"), ',')
Output: ["a" "bb" "c"]
```

**Example 2:**

```
Input:  Fields(nil, []byte("a,,b"), ',')
Output: ["a" "" "b"]
```

_Explanation:_ The empty middle field is a field.

**Example 3:**

```
Input:  appending to field 0
Output: does not touch field 1
```

_Explanation:_ Each view is capacity-capped.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Views instead of copies** | The fields point into the caller's line. |
| 2 | **Three-index slicing** | `line[start:i:i]` stops an append from spilling into the next field. |
| 3 | **Append-style APIs** | A `dst` parameter lets the caller own the header slice across lines. |
| 4 | **Virtual trailing separator** | Running to `len(line)` inclusive closes the last field. |

## Hint

Two things per field: where it starts and ends, and what its capacity should be.

## Validate

```bash
make verify
```

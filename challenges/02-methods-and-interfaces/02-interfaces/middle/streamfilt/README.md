# Stream Filter

**Level:** middle
**Topic:** 02-methods-and-interfaces → 02-interfaces

## Context

A log pipeline filters lines as they stream past, without buffering the whole file.

## Task

Implement the stub(s) in [streamfilt.go](streamfilt.go):

1. Implement `Match` on `Contains` and `MinLen`.
2. Implement `Not`, which inverts any predicate.
3. Implement `FilterStream`, which reads from a `Source` and returns the matching lines, using one pass and at most one buffer of results.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Contains{Sub: "err"}.Match("error")
Output: true
```

**Example 2:**

```
Input:  Not{Inner: Contains{Sub: "err"}}.Match("error")
Output: false
```

**Example 3:**

```
Input:  FilterStream(src, MinLen{N: 3})
Output: only lines of 3+ bytes
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Predicate interface** | Filters compose without touching the pipeline. |
| 2 | **Streaming** | Efficiency: read one line at a time, never the whole source. |
| 3 | **Wrapping predicates** | `Not` decorates any predicate, including another `Not`. |

## Hint

Loop `src.Next()` until `ok` is false, appending only the matches.

## Validate

```bash
make verify
```

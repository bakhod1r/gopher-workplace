# Join Without The Intermediate Strings

**Level:** junior
**Topic:** 10-advanced-topics / 01-memory-management-in-depth

## Context

A CSV header is assembled with `line += col + ","`. With 200 columns that is 400 short-lived strings per row, and the rows are the whole file.

## Task

Implement [builder.go](builder.go):

1. Concatenate `parts` with `sep` between them.
2. Size the buffer once before writing — at most 3 allocations for a 64-part join.
3. An empty input returns the empty string.

Replace the stub body in [builder.go](builder.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Join([]string{"a","b","c"}, "-")
Output: "a-b-c"
```

**Example 2:**

```
Input:  Join([]string{"solo"}, ",")
Output: "solo"
```

_Explanation:_ No separator with a single part.

**Example 3:**

```
Input:  Join(nil, ",")
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **String immutability** | `+=` cannot extend a string; it allocates a new one and copies both sides. |
| 2 | **strings.Builder** | A growable byte buffer that hands out its bytes as a string without a final copy. |
| 3 | **Grow** | One `Grow(n)` replaces every doubling step. |

## Hint

Compute the exact final length before you write the first byte.

## Validate

```bash
make verify
```

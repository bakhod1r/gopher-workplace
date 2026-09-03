# The Concatenation That Copied Everything Each Time

**Level:** middle
**Topic:** 10-advanced-topics / 10-advanced-topics

## Context

A report builder concatenates a few hundred fragments with `+=`. The fragments are short, the report is not, and the function dominates the profile.

## Task

Fix the single planted bug in [concatloop.go](concatloop.go):

1. Concatenate the parts end to end.
2. Fix the single bug so the work is linear: at most a handful of allocations for 128 parts.
3. An empty input returns the empty string.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Join([]string{"a","bc","d"})
Output: "abcd"
```

**Example 2:**

```
Input:  128 parts
Output: at most 3 allocations
```

_Explanation:_ Not 128.

**Example 3:**

```
Input:  Join(nil)
Output: ""
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **String immutability** | `+=` cannot extend a string; it builds a new one and copies both sides. |
| 2 | **Quadratic accumulation** | Joining n parts copies O(n²) bytes. |
| 3 | **strings.Builder** | One growing buffer, handed out as a string without a final copy. |
| 4 | **Grow** | Sizing it once removes even the doubling. |

## Hint

The loop is fine. What it accumulates into is not.

## Validate

```bash
make verify
```

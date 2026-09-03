# The Function Entry Point Blamed On Its Neighbour

**Level:** senior  
**Topic:** 11-performance-engineering

## Context

Short, hot, frequently-called functions show up attributed to whatever happens to be laid out before them in the binary. Long functions look fine. The pattern is the tell: the samples that land exactly on a function's first instruction are being resolved one symbol too low.

## Task

Fix the single planted bug in [symbolizeoffbug.go](symbolizeoffbug.go):

1. Find and fix the one bug so an address landing exactly on a symbol's start resolves to that symbol.
2. Addresses inside a function must keep resolving to it.
3. An address below the first symbol must still report `"", false`.

Change only the code between the `CHANGE CODE` markers.

## Examples

**Example 1:**

```
Input:  Resolve([{100 a} {200 b}], 200)
Output: "b", true
```

**Example 2:**

```
Input:  Resolve([{100 a} {200 b}], 150)
Output: "a", true
```

**Example 3:**

```
Input:  Resolve([{100 a}], 99)
Output: "", false
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **`sort.Search` returns the first true** | The predicate defines the boundary, and `>=` versus `>` moves it by one symbol. |
| 2 | **Half-open ranges** | A symbol owns `[start, nextStart)`, so its own start belongs to it. |
| 3 | **Entry points are heavily sampled** | Short functions spend a large fraction of their time on their first instruction. |

## Hint

The predicate should find the first symbol that starts strictly *after* the address.

## Validate

```bash
make verify
```

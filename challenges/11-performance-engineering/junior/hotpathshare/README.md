# How Much Of The Profile Is This?

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

Before optimising anything, answer one question: what share of the total does this code path own? Amdahl's law makes the answer decisive — a 10x speedup of a 2% path buys you 1.8%.

## Task

Implement `Share` in [hotpathshare.go](hotpathshare.go):

1. Sum the flat values of the named functions and divide by the profile total.
2. Names absent from `flat` contribute nothing; a name listed twice counts once.
3. An empty profile returns `0`.

## Examples

**Example 1:**

```
Input:  Share({a:3 b:1}, [a])
Output: 0.75
```

**Example 2:**

```
Input:  Share({a:3 b:1}, [a a a])
Output: 0.75
```

**Example 3:**

```
Input:  Share({a:3 b:1}, [a nope])
Output: 0.75
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Amdahl's law in practice** | The share bounds every possible win from optimising that path. |
| 2 | **Deduplicate the selection** | Double-counting a name can push the share above 1 and hide the real ceiling. |
| 3 | **Two passes, one map** | Total over the whole map, selected sum over the deduplicated names. |

## Topics used again

Maps as sets, float division, guards.

## Hint

Deduplicate the names into a set first, then walk the map once and accumulate both sums.

## Validate

```bash
make verify
```

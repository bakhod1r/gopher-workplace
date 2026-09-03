# The Flat Column

**Level:** junior  
**Topic:** 11-performance-engineering

## Context

`go tool pprof -top` opens with the flat column: time the CPU spent *inside* each function, ignoring whatever that function called. It is the first thing anyone reads in a profile, and it is nothing more than a grouped sum over samples.

## Task

Implement `FlatSum` in [flatsum.go](flatsum.go):

1. Total `Self` per function name.
2. Skip samples whose `Self` is zero or negative.
3. A function left with no samples must not appear as a key; an empty input returns an empty, non-nil map.

## Examples

**Example 1:**

```
Input:  [{a 3} {b 1} {a 2}]
Output: {a:5 b:1}
```

**Example 2:**

```
Input:  [{a 5} {b 0} {c -3}]
Output: {a:5}
```

**Example 3:**

```
Input:  nil
Output: {} (non-nil)
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Flat is self time** | It excludes callees, which is why a hot leaf function stands out. |
| 2 | **A profile is just samples** | Every pprof column is an aggregation over a sample list. |
| 3 | **Absent vs zero** | A key with value `0` claims the function was sampled; leaving it out says it was not. |

## Topics used again

Maps, structs, `range`.

## Hint

`make(map[string]int64)` first, then one `range` with a guard.

## Validate

```bash
make verify
```

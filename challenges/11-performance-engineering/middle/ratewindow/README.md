# Rates Over A Time Window

**Level:** middle  
**Topic:** 11-performance-engineering

## Context

"How many requests per second" is always a question about a window, and the window's width is what makes the number mean anything: the same three events are 3 rps over a second and 6 rps over half of one. Getting the boundaries wrong — counting an event in two adjacent windows — makes a graph that adds up to more traffic than the system ever saw.

## Task

Implement all three in [ratewindow.go](ratewindow.go):

1. `CountIn` counts the events in the half-open window `[from, from+width)` using binary searches over the sorted timestamps.
2. `RatePerSec` divides that count by the width in seconds; a non-positive width gives `0`.
3. `SumIn` totals the events' `Value` in the same window.

## Examples

**Example 1:**

```
Input:  CountIn(events at 0,50,99,100,250, 0, 100)
Output: 3
```

**Example 2:**

```
Input:  CountIn(the same, 100, 100)
Output: 1
```

**Example 3:**

```
Input:  three events in the first second; RatePerSec(_, 0, 0.5s)
Output: 6
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Half-open windows tile exactly** | `[a,b)` and `[b,c)` cover everything once, with no event double-counted. |
| 2 | **The rate depends on the width** | A rate without its window is not a number, it is a rumour. |
| 3 | **Binary search on sorted timestamps** | Two searches bound the window in log time regardless of series length. |

## Topics used again

Binary search, half-open intervals, int64 time arithmetic.

## Hint

Find the first index at or after `from` and the first at or after `from+width`; the difference is the count.

## Validate

```bash
make verify
```

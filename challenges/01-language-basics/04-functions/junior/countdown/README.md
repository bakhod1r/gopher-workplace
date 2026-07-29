# Countdown Slice

**Level:** junior
**Topic:** 01-language-basics → 04-functions · _loops_

## Context

A classic `for` loop with a decrementing counter builds a descending sequence.

## Task

Implement `Countdown` in [countdown.go](countdown.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Countdown(3)
Output: [3 2 1]
```

**Example 2:**

```
Input:  Countdown(1)
Output: [1]
```

**Example 3:**

```
Input:  Countdown(0)
Output: []
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Three-clause for** | `for i := n; i >= 1; i--` counts down. |
| 2 | **append** | Grow the result each iteration. |
| 3 | **Empty case** | `n <= 0` never enters the loop. |

## Hint

Loop `for i := n; i >= 1; i-- { out = append(out, i) }`.

## Validate

```bash
make verify
```

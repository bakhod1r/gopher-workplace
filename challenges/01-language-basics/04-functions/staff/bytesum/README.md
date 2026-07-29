# Accumulator Type Overflow

**Level:** staff
**Topic:** 01-language-basics → 04-functions · _conditionals_

## Context

A `uint8` accumulator wraps modulo 256, so 200+100+50 = 350 becomes 94. The
accumulator's type must be wide enough for the total, not just for each element.

## Task

Fix the accumulator type in [bytesum.go](bytesum.go).

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  Sum([]byte{200, 100, 50})
Output: 350
```

**Example 2:**

```
Input:  Sum([]byte{255, 1})
Output: 256
```

**Example 3:**

```
Input:  Sum(nil)
Output: 0
```

## Topics to Master

Only concepts taught at or before this slot (scope rule, see GENERATION.md).

| # | Topic | What to understand |
|---|-------|--------------------|
| 1 | **Accumulator width** | It must hold the sum, not just each term. |
| 2 | **Unsigned wraparound** | uint8 arithmetic is mod 256. |
| 3 | **Type of the running total** | Use int (or a wide unsigned). |

## Hint

Accumulate into an `int`: `var total int`, add `int(b)`, return `total`.

## Validate

```bash
make verify
```

# A Scratch Array That Never Leaves

**Level:** junior
**Topic:** 10-advanced-topics / 02-escape-analysis

## Context

A logging helper counts digits by rendering the number with `strconv.Itoa` and taking the length. One allocation per log line, on the hottest path in the service.

## Task

Implement [localarray.go](localarray.go):

1. Return how many decimal digits `n` has, ignoring the sign.
2. `Digits(0)` is 1.
3. Use a fixed-size local array; the function must allocate nothing.

Replace the stub body in [localarray.go](localarray.go) with a working implementation.

## Examples

**Example 1:**

```
Input:  Digits(1234)
Output: 4
```

**Example 2:**

```
Input:  Digits(0)
Output: 1
```

_Explanation:_ Zero is one digit.

**Example 3:**

```
Input:  Digits(-99)
Output: 2
```

_Explanation:_ The sign does not count.

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Fixed-size locals** | `var buf [20]byte` is a stack object when it does not escape. |
| 2 | **Arrays vs slices** | An array has its size in its type, so the compiler can place it in the frame. |
| 3 | **Avoiding conversions** | `strconv.Itoa` allocates a string you throw away. |

## Hint

Twenty bytes is enough for any int64. Where should those twenty bytes live?

## Validate

```bash
make verify
```

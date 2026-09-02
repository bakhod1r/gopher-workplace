# Aisle Seats

**Level:** junior
**Topic:** 06-concurrency → 02-channels

## Context

The ticketing system's seat-reservation feed streams seat numbers for a
row. The aisle-side seats are the even-numbered ones, and only those are
forwarded to the accessibility allocator.

## Task

Implement `AisleSeats` in [seatfilter.go](seatfilter.go) so that:

1. A goroutine walks `seats` and sends only even-numbered seats on the channel.
2. It closes the channel when the row is exhausted.
3. The allocator collects them in order; a row with no aisle seats gives an empty, non-nil slice.

Do **not** change the function signature or the tests.

## Examples

**Example 1:**

```
Input:  AisleSeats([]int{1, 2, 3, 4})
Output: [2 4]
```

**Example 2:**

```
Input:  AisleSeats([]int{1, 3})
Output: []
```

**Example 3:**

```
Input:  AisleSeats(nil)
Output: []
```

## Topics to Master

| # | Topic | What to understand |
|---|-------|---------------------|
| 1 | **Conditional send** | Not every input produces an output value. |
| 2 | **Close is unconditional** | The channel closes even when no seat matched. |
| 3 | **`seat%2 == 0`** | Works for negatives too: `-4 % 2 == 0`. |

## Hint

`close(aisle)` goes after the loop, outside the `if` — it must run even
when no seat matched.

## Validate

```bash
make verify
```

# Gap size in two-pointer windows

## The idea

The lead pointer must be exactly n ahead; an inclusive bound makes it n+1 ahead and shifts the result.

## Why it matters

Off-by-one gaps are the classic nth-from-end bug.

## Watch out

- `i <= n` advances n+1 times.
- The gap equals the number of lead-advance steps.

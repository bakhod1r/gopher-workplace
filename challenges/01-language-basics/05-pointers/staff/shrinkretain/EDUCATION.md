# Clearing dropped pointer slots

## The idea

Length reduction leaves stale pointers in the backing array; for pointer/interface elements you must nil the vacated slots or the referents leak.

## Why it matters

The 'nil out on pop' pattern is standard in Go container code to avoid memory leaks.

## Watch out

- `s[:last]` keeps the old element in the array's spare capacity.
- Set `s[last] = nil` before shrinking for pointer element types.

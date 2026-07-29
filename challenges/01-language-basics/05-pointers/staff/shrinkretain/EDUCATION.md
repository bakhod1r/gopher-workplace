# Clearing dropped pointer slots

## Intuition

Length reduction leaves stale pointers in the backing array; for pointer/interface elements you must nil the vacated slots or the referents leak.

## Approach

1. Just reslicing `s[:last]` leaves the popped pointer live in the backing array, leaking it.
2. Nil the slot first: `s[last] = nil`, then `*sp = s[:last]`.

## Solution

```go
func Pop(sp *[]*int) *int {
	s := *sp
	last := len(s) - 1
	v := s[last]
	s[last] = nil
	*sp = s[:last]
	return v
}
```

## Walkthrough

The bug keeps the last element reachable through the array's capacity, preventing collection. Nilling the slot before shrinking releases it.

## Pitfalls

- `s[:last]` keeps the old element in the array's spare capacity.
- Set `s[last] = nil` before shrinking for pointer element types.

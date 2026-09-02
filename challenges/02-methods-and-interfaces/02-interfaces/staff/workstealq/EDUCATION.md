# Work-Stealing Deque

## Intuition

Owner and thieves work opposite ends so they almost never touch the same element — that is the whole point of the design. The interesting case is the last item, where both ends converge and exactly one side must win.

## Approach

1. `Push` appends and advances `bottom`.
2. `Pop` returns false when `bottom <= top`, otherwise decrements `bottom` and reads that slot.
3. `Steal` returns false when `top >= bottom`, otherwise reads `top` and advances it.
4. The index invariant `top <= bottom` is what makes the contended last item resolvable.

## Solution

```go
func (d *Deque) Pop() (int, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.bottom <= d.top {
		return 0, false
	}
	d.bottom--
	v := d.items[d.bottom]
	return v, true
}

func (d *Deque) Steal() (int, bool) {
	d.mu.Lock()
	defer d.mu.Unlock()

	if d.top >= d.bottom {
		return 0, false
	}
	v := d.items[d.top]
	d.top++
	return v, true
}
```

## Walkthrough

`TestRaceOnLastItem` runs 200 trials of a `Pop` against a `Steal` over a single element. The `top`/`bottom` comparison inside the critical section is what guarantees exactly one of them succeeds.

## Pitfalls

- Checking emptiness with `len(d.items)` instead of `bottom - top` — stolen slots are still in the slice.
- Letting `Pop` and `Steal` both take the last element because the bounds check happens outside the critical section.
- In a genuinely lock-free version, forgetting the CAS on `top` for the last item — the classic Chase-Lev bug this locked version sidesteps.

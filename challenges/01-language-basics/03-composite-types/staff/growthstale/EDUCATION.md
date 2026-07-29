# append can invalidate pointers

## Intuition

When `append` exceeds capacity, it allocates a **new** backing array and copies.
Any pointer or sub-slice into the old array still references the old memory:

```go
p := &s[0]
s = append(s, x) // may reallocate
s[0] = 99        // correct: writes the current array (don't use *p)
```

## Approach

1. Bug: after append reallocates, `p` points into the OLD array; `*p = 99` writes to freed/detached memory, leaving s[0]=10.
2. Fix: write through the live slice `s[0] = 99` so the current backing array is updated.

## Solution

```go
func BumpFirst(val int) (int, []int) {
	s := make([]int, 1, 1) // len 1, cap 1 (full)
	s[0] = 10
	p := &s[0]
	s = append(s, val) // reallocates; p now points at the OLD array
	_ = p
	s[0] = 99
	return s[0], s
}
```

## Walkthrough

s=[10] cap1. append(s,7) allocates a new array [10 7]; p still points at the old [10]. *p=99 mutates the old array (invisible). s[0]=99 mutates the new array -> returns 99, [99 7].

## Pitfalls

- Growth is capacity-triggered; with spare cap, `p` would still be valid — making
  the bug intermittent.
- Sub-slices taken before an append can likewise detach.
- Prefer indices over long-lived element pointers into growable slices.

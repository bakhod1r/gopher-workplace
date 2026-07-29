# Order-preserving dedup

## Intuition

Keep a set of already-emitted values; emit an element the **first** time you see
it (when it is NOT in the set):

```go
if _, ok := seen[x]; !ok {
	out = append(out, x); seen[x] = struct{}{}
}
```

## Approach

1. Bug: if ok { appends only when the value was ALREADY seen, inverting the logic (and it never records first sightings correctly). 2. Fix: if !ok { so we append on first occurrence and then mark it seen. 3. The seen map preserves first-seen order via the single append per new value.

## Solution

```go
func Unique(xs []int) []int {
	seen := make(map[int]struct{})
	out := []int{}
	for _, x := range xs {
		_, ok := seen[x]
		if !ok {
			out = append(out, x)
			seen[x] = struct{}{}
		}
	}
	return out
}
```

## Walkthrough

[1,1,2]: x=1 ok=false -> append 1, mark. x=1 ok=true -> skip. x=2 ok=false -> append 2. Result [1,2]. With the bug, first 1 (ok=false) is skipped, breaking output.

## Pitfalls

- Emit on absence, then record; recording first would suppress everything.
- The map gives O(1) membership; a slice scan would be O(n²).
- `slices.Compact` only removes *consecutive* dups — different semantics.

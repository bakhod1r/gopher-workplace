# The merge step

## Intuition

Two indices walk the sorted inputs; append the smaller front element, advance
that side, and finally drain whatever remains:

```go
i, j := 0, 0
for i < len(a) && j < len(b) {
	if a[i] <= b[j] { out = append(out, a[i]); i++ } else { out = append(out, b[j]); j++ }
}
out = append(out, a[i:]...)
out = append(out, b[j:]...)
```

## Approach

1. Two pointers i,j at heads of a,b.
2. While both remain, append the smaller (a[i]<=b[j] keeps stability).
3. Append the remaining tail of whichever slice is left.
4. Return merged slice.

## Solution

```go
func Merge(a, b []int) []int {
	out := make([]int, 0, len(a)+len(b))
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}
	out = append(out, a[i:]...)
	out = append(out, b[j:]...)
	return out
}
```

## Walkthrough

a=[1,3,5] b=[2,4,6]: 1<=2 take1; 3>2 take2; 3<=4 take3; 5>4 take4; 5<=6 take5; a done, append b tail [6] -> [1,2,3,4,5,6].

## Pitfalls

- Use `<=` (not `<`) to keep the merge stable for equal keys.
- Remember to drain both tails — one will be empty.
- Pre-size the output to `len(a)+len(b)` to avoid regrowth.

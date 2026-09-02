# Merge Sorted

## Intuition

Merging needs only two questions per step: what is next on each side, and give me that one. `Peek`/`Next` is exactly that contract, and it keeps the merge O(n+m).

## Approach

1. `Peek` returns `Data[pos]` without advancing; `Next` reuses `Peek` and then advances.
2. In `Merge`, peek both feeds each iteration.
3. Return when both are drained; otherwise consume from the side with the smaller head (`<=` keeps it stable).
4. When one side is drained, the other's condition takes over and drains it.

## Solution

```go
func (f *SliceFeed) Peek() (int, bool) {
	if f.pos >= len(f.Data) {
		return 0, false
	}
	return f.Data[f.pos], true
}

func (f *SliceFeed) Next() (int, bool) {
	v, ok := f.Peek()
	if !ok {
		return 0, false
	}
	f.pos++
	return v, true
}

func Merge(a, b Feed) []int {
	var out []int
	for {
		av, aok := a.Peek()
		bv, bok := b.Peek()
		switch {
		case !aok && !bok:
			return out
		case aok && (!bok || av <= bv):
			v, _ := a.Next()
			out = append(out, v)
		default:
			v, _ := b.Next()
			out = append(out, v)
		}
	}
}
```

## Walkthrough

`[1 1]` and `[1]`: `av <= bv` keeps taking from `a` while the heads tie, then `a` drains and `b`'s single element follows — three elements, order preserved.

## Pitfalls

- `av < bv` instead of `<=` — still correct output here, but it flips the relative order of equal elements between feeds.
- Consuming with `Next` while comparing, which drops a value when the other side wins.
- Concatenating and calling `sort.Ints`, which throws away the fact that the inputs are already sorted.

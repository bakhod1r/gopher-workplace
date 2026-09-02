# sort.Interface

## Intuition

`sort.Sort` knows nothing about players. It moves elements through `Swap` and asks `Less` for the order, so the ordering policy lives entirely in your three methods.

## Approach

1. `Len` and `Swap` are mechanical.
2. In `Less`, compare scores with `>` for descending order.
3. When scores are equal, fall through to `b[i].Name < b[j].Name`.
4. In `TopN`, sort, clamp `n` to the length, then collect names.

## Solution

```go
func (b ByScore) Len() int { return len(b) }

func (b ByScore) Less(i, j int) bool {
	if b[i].Score != b[j].Score {
		return b[i].Score > b[j].Score
	}
	return b[i].Name < b[j].Name
}

func (b ByScore) Swap(i, j int) { b[i], b[j] = b[j], b[i] }

func TopN(players []Player, n int) []string {
	sort.Sort(ByScore(players))
	if n > len(players) {
		n = len(players)
	}
	out := make([]string, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, players[i].Name)
	}
	return out
}
```

## Walkthrough

With `ann` and `bob` both on 5, the score branch is skipped and the name comparison puts `ann` first — a deterministic tie-break the test relies on.

## Pitfalls

- Returning `b[i].Score < b[j].Score`, which sorts ascending.
- Comparing names first, which ignores the score entirely.
- Indexing `players[i]` past the end when `n` exceeds the slice length.

# The Format Flame Graphs Eat

## Intuition

Folding is a group-by over the whole call path. Turning the path into a string gives you a comparable, hashable key and the output format at the same time.

## Approach

1. Join each valid stack with `;` and accumulate into a map.
2. Move the pairs into a slice and sort by value descending, then key ascending.
3. Format each pair.

## Solution

```go
func Fold(samples []Sample) []string {
	totals := make(map[string]int64)
	for _, s := range samples {
		if s.Value <= 0 || len(s.Stack) == 0 {
			continue
		}
		totals[strings.Join(s.Stack, ";")] += s.Value
	}
	type row struct {
		key string
		val int64
	}
	rows := make([]row, 0, len(totals))
	for k, v := range totals {
		rows = append(rows, row{k, v})
	}
	slices.SortFunc(rows, func(a, b row) int {
		if c := cmp.Compare(b.val, a.val); c != 0 {
			return c
		}
		return cmp.Compare(a.key, b.key)
	})
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, fmt.Sprintf("%s %d", r.key, r.val))
	}
	return out
}
```

## Walkthrough

`a;b` and `b;a` hash differently, which is exactly right: they are two distinct paths through the program that happen to involve the same two functions.

## Pitfalls

- Aggregating by leaf function instead of by whole stack, which collapses paths a flame graph needs to keep apart.
- Sorting on value only, so equal-valued stacks shuffle between runs.
- Using a separator that can appear in a symbol name, silently merging two stacks.

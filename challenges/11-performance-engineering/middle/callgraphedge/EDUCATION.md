# Turning Stacks Into A Graph

## Intuition

A stack is a path. A path of `n` nodes contains `n-1` edges, and summing those edges over every sample gives the weighted graph.

## Approach

1. For each valid sample, walk adjacent pairs and accumulate into a map keyed by a `{caller, callee}` struct.
2. Flatten and sort with three keys.
3. `CalleesOf` filters the arcs and sorts them.

## Solution

```go
type pair struct{ caller, callee string }

func Edges(samples []Sample) []Edge {
	totals := make(map[pair]int64)
	for _, s := range samples {
		if s.Value <= 0 || len(s.Stack) < 2 {
			continue
		}
		for i := 0; i+1 < len(s.Stack); i++ {
			totals[pair{s.Stack[i], s.Stack[i+1]}] += s.Value
		}
	}
	out := make([]Edge, 0, len(totals))
	for p, v := range totals {
		out = append(out, Edge{p.caller, p.callee, v})
	}
	slices.SortFunc(out, func(a, b Edge) int {
		if c := cmp.Compare(b.Value, a.Value); c != 0 {
			return c
		}
		if c := cmp.Compare(a.Caller, b.Caller); c != 0 {
			return c
		}
		return cmp.Compare(a.Callee, b.Callee)
	})
	return out
}

func CalleesOf(edges []Edge, caller string) []string {
	type row struct {
		name string
		val  int64
	}
	rows := make([]row, 0, len(edges))
	for _, e := range edges {
		if e.Caller == caller {
			rows = append(rows, row{e.Callee, e.Value})
		}
	}
	slices.SortFunc(rows, func(a, b row) int {
		if c := cmp.Compare(b.val, a.val); c != 0 {
			return c
		}
		return cmp.Compare(a.name, b.name)
	})
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, r.name)
	}
	return out
}
```

## Walkthrough

`[rec rec rec]` contains two adjacent `rec→rec` pairs, so a value of 2 contributes 4 to that self-arc — direct recursion genuinely traversed the edge twice.

## Pitfalls

- Deduplicating adjacent pairs the way cum deduplicates frames; edges count traversals, not membership.
- Emitting an arc for every frame pair rather than adjacent ones, inventing calls that never happened.
- Sorting by value alone and shipping a graph whose arc order changes every run.

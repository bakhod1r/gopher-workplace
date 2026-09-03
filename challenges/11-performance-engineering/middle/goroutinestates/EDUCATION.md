# Reading A Goroutine Dump

## Intuition

A dump is a list of goroutines. Everything useful comes from grouping it: by state to see the shape of the hang, by state and site to find where they all piled up.

## Approach

1. `Count` groups by state, normalising the empty one.
2. `Blocked` counts everything outside the two running states.
3. `LeakSuspects` groups by `{state, site}`, drops running states, filters by threshold, sorts.

## Solution

```go
func Count(gs []G) map[string]int {
	out := make(map[string]int)
	for _, g := range gs {
		state := g.State
		if state == "" {
			state = "unknown"
		}
		out[state]++
	}
	return out
}

func running(state string) bool { return state == "running" || state == "runnable" }

func Blocked(gs []G) int {
	n := 0
	for _, g := range gs {
		if !running(g.State) {
			n++
		}
	}
	return n
}

func LeakSuspects(gs []G, threshold int) []string {
	if threshold < 1 {
		threshold = 1
	}
	type key struct{ state, site string }
	counts := make(map[key]int)
	for _, g := range gs {
		if running(g.State) {
			continue
		}
		counts[key{g.State, g.Top}]++
	}
	type row struct {
		key key
		n   int
	}
	rows := make([]row, 0, len(counts))
	for k, n := range counts {
		if n >= threshold {
			rows = append(rows, row{k, n})
		}
	}
	slices.SortFunc(rows, func(a, b row) int {
		if c := cmp.Compare(b.n, a.n); c != 0 {
			return c
		}
		if c := cmp.Compare(a.key.state, b.key.state); c != 0 {
			return c
		}
		return cmp.Compare(a.key.site, b.key.site)
	})
	out := make([]string, 0, len(rows))
	for _, r := range rows {
		out = append(out, r.key.site)
	}
	return out
}
```

## Walkthrough

Grouping by state *and* site rather than site alone matters when the same function blocks in two different ways — a `select` and a `chan receive` at one line are two distinct situations.

## Pitfalls

- Counting `runnable` as blocked, which turns a busy scheduler into a phantom leak.
- Reporting sites without a threshold, so ordinary waiting goroutines drown the real signal.
- Dropping the state from the grouping key and merging two unrelated wait conditions.

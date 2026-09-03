# Both Columns, One Pass

## Intuition

Each sample says two things at once: this leaf was executing, and this whole chain was responsible. Record both while you have the stack in hand.

## Approach

1. For each valid sample, credit the leaf's `Flat`.
2. Walk the stack with a seen-set, crediting each new frame's `Cum`.
3. `Leaves` filters and sorts.

## Solution

```go
func Analyze(samples []Sample) map[string]Node {
	out := make(map[string]Node)
	for _, s := range samples {
		if s.Value <= 0 || len(s.Stack) == 0 {
			continue
		}
		leaf := s.Stack[len(s.Stack)-1]
		n := out[leaf]
		n.Flat += s.Value
		out[leaf] = n
		seen := make(map[string]bool, len(s.Stack))
		for _, fn := range s.Stack {
			if seen[fn] {
				continue
			}
			seen[fn] = true
			n := out[fn]
			n.Cum += s.Value
			out[fn] = n
		}
	}
	return out
}

func Leaves(nodes map[string]Node, minFlat int64) []string {
	out := make([]string, 0, len(nodes))
	for fn, n := range nodes {
		if n.Flat >= minFlat {
			out = append(out, fn)
		}
	}
	slices.SortFunc(out, func(a, b string) int {
		if c := cmp.Compare(nodes[b].Flat, nodes[a].Flat); c != 0 {
			return c
		}
		return cmp.Compare(a, b)
	})
	return out
}
```

## Walkthrough

`out[fn].Cum += v` does not compile: map values are not addressable. The read-modify-write through a local `Node` is the idiomatic way around it.

## Pitfalls

- Crediting `Flat` to every frame, which turns the flat column into a second cum column.
- Sharing one seen-set across samples, so each function gets cum credit only once for the whole profile.
- Sorting by flat alone and shipping a report that reorders itself.

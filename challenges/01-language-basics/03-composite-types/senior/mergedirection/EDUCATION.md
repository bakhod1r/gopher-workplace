# Override order in a merge

## Intuition

When copying two maps into one, the **last** write to a key wins. For "override
on top of base", copy the base first, then the overrides:

```go
for k, v := range base { out[k] = v }
for k, v := range over { out[k] = v } // over wins
```

## Approach

1. Bug: over is copied first, then base overwrites it, so base wins on key collisions (wrong). 2. Later writes to the same key win in a map. 3. Fix: copy base first, then over on top, so over wins.

## Solution

```go
func Merge(base, over map[string]int) map[string]int {
	out := make(map[string]int)
	for k, v := range base {
		out[k] = v
	}
	for k, v := range over {
		out[k] = v
	}
	return out
}
```

## Walkthrough

base={b:2}, over={b:9}. Buggy order: out[b]=9 then out[b]=2 -> 2. Correct order: out[b]=2 then out[b]=9 -> 9.

## Pitfalls

- Order of the two loops is the whole behavior.
- Within one map, iteration order is random but irrelevant here.
- Document precedence explicitly; it's easy to get backwards.

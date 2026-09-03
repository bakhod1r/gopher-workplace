# Many Machines, One Profile

## Intuition

Merging profiles is addition on two axes at once: per-function values and the sample count that gives them scale.

## Approach

1. Allocate a fresh result map.
2. Add every profile's entries and sample count into it.

## Solution

```go
func Merge(profiles []Profile) Profile {
	out := Profile{Flat: make(map[string]int64)}
	for _, p := range profiles {
		out.Samples += p.Samples
		for fn, v := range p.Flat {
			out.Flat[fn] += v
		}
	}
	return out
}

func Total(p Profile) int64 {
	var total int64
	for _, v := range p.Flat {
		total += v
	}
	return total
}
```

## Walkthrough

Ranging a nil map iterates zero times, so an empty profile needs no special case — only its sample count still counts.

## Pitfalls

- Using the first profile's map as the accumulator, which mutates the caller's data.
- Merging values while forgetting the sample counts, making the result impossible to normalise.
- Merging profiles taken over different durations without scaling them first.

# Iterator Chain

## Intuition

Pull-based iterators invert the usual pipeline: nothing happens until the consumer asks, and each stage holds one element at a time instead of a whole intermediate slice.

## Approach

1. `MapIter.Next` pulls one element and applies `Fn`.
2. `FilterIter.Next` loops, discarding non-matching elements, until it finds one or the source drains.
3. `Collect` drains the outermost iterator into a slice.
4. Building the structs performs no reads at all.

## Solution

```go
func (m *MapIter) Next() (int, bool) {
	v, ok := m.Inner.Next()
	if !ok {
		return 0, false
	}
	return m.Fn(v), true
}

func (f *FilterIter) Next() (int, bool) {
	for {
		v, ok := f.Inner.Next()
		if !ok {
			return 0, false
		}
		if f.Pred(v) {
			return v, true
		}
	}
}

func Collect(it Iter) []int {
	var out []int
	for {
		v, ok := it.Next()
		if !ok {
			return out
		}
		out = append(out, v)
	}
}
```

## Walkthrough

`TestEachElementReadOnce` builds map → filter → map over five elements and asserts `Reads == 5`. Any stage that materialised its input would read the source more than once or read it eagerly.

## Pitfalls

- `FilterIter.Next` returning `(0, false)` on the first non-match, which truncates the stream.
- Collecting inside a stage to "simplify", which reintroduces the intermediate slices.
- Applying `Fn` before checking `ok`, transforming a zero value that was never yielded.

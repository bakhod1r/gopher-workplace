# Sorted Keys Of Any String-Keyed Map

## Intuition

Map iteration order in Go is deliberately randomised, so any reflective dump must sort. The key type is part of the map's type, so you can validate the shape before reading a single entry.

## Approach

1. Reject anything that is not a map with a string key type.
2. Preallocate to `rv.Len()` and collect `k.String()` for each key.
3. Sort and return.

## Solution

```go
import (
	"reflect"
	"sort"
)

// Keys returns the keys of m sorted in ascending order.
//
// m must be a map with string keys; anything else yields nil. The value
// type does not matter.
//
// Examples:
//
// 	Keys(map[string]int{"b": 1, "a": 2}) => []string{"a", "b"}
func Keys(m any) []string {
	rv := reflect.ValueOf(m)
	if rv.Kind() != reflect.Map || rv.Type().Key().Kind() != reflect.String {
		return nil
	}
	out := make([]string, 0, rv.Len())
	for _, k := range rv.MapKeys() {
		out = append(out, k.String())
	}
	sort.Strings(out)
	return out
}
```

## Walkthrough

`Keys` on a four-entry map collects the keys in whatever order the runtime offers, then sorts them — so fifty calls produce fifty identical slices.

## Pitfalls

- `reflect.ValueOf(nil).Kind()` is invalid, which the map check already rejects.
- Using `fmt.Sprint(k)` instead of `k.String()`; the kind check has already guaranteed a string.

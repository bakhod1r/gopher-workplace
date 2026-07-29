# Distinct allocations per iteration

## Intuition

Taking the address of a single hoisted variable shares one allocation; declaring the variable inside the loop gives each pointer its own heap object.

## Approach

1. A single hoisted `it` is reused, so all pointers alias the last value.
2. Declare `it := Item{V: v}` inside the loop for fresh storage.

## Solution

```go
type Item struct{ V int }

func Items(vs []int) []*Item {
	var out []*Item
	for _, v := range vs {
		it := Item{V: v}
		out = append(out, &it)
	}
	return out
}
```

## Walkthrough

With the shared `it`, every `&it` points at one struct holding the final value. A per-iteration `it` gives each pointer its own item.

## Pitfalls

- `&it` of a hoisted `it` aliases one struct.
- Declare `it := Item{...}` inside the loop for distinct pointers.

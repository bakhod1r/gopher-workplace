# Sorting records

## Intuition

`sort.Slice` takes a "less" function. For a stable, deterministic order, compare
the primary field, then break ties on a secondary field:

```go
c := append([]Person{}, people...)
sort.Slice(c, func(i, j int) bool {
	if c[i].Age != c[j].Age { return c[i].Age < c[j].Age }
	return c[i].Name < c[j].Name
})
```

## Approach

1. Copy input to a new slice (input must not mutate).
2. sort.Slice with comparator: Age ascending; if Age equal, Name ascending.
3. Return the copy.

## Solution

```go
import "sort"

type Person struct {
	Name string
	Age  int
}

func ByAge(people []Person) []Person {
	out := make([]Person, len(people))
	copy(out, people)
	sort.Slice(out, func(i, j int) bool {
		if out[i].Age != out[j].Age {
			return out[i].Age < out[j].Age
		}
		return out[i].Name < out[j].Name
	})
	return out
}
```

## Walkthrough

bob(30),amy(25),cid(30): amy has lowest age; bob and cid tie at 30, bob<cid by name -> [amy,bob,cid].

## Pitfalls

- `sort.Slice` is **not** stable; encode tie-breaks in the comparator (or use
  `sort.SliceStable`).
- Sort mutates in place — copy if the input must survive.
- The comparator must be a strict weak ordering (no `<=`).

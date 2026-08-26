# Method Expressions in sort.Slice

## Intuition

`sort.Slice` wants `func(i, j int) bool` — index-based. But our comparison is
value-based: `func(Person, Person) bool`. The bridge is a closure that looks up
elements by index.

Method expressions like `Person.ByAge` keep the code clean: the caller names the
sort key, the function does the mechanical sorting.

## Approach

1. Call `sort.Slice` with a closure that applies `less` to the indexed elements.

## Solution

```go
func SortByField(people []Person, less func(Person, Person) bool) {
	sort.Slice(people, func(i, j int) bool {
		return less(people[i], people[j])
	})
}
```

## Walkthrough

`SortByField([Charlie:30, Alice:20], Person.ByAge)`:
- `less = Person.ByAge`.
- Closure: `less(people[0], people[1])` → `30 < 20` → false → swap.
- Result: `[Alice:20, Charlie:30]`.

## Pitfalls

- Passing `less` directly to `sort.Slice` won't compile — signatures don't match.
- Method expression `Person.ByAge` has the receiver as the first arg; method
  value `p.ByAge` would be `func(Person) bool` — wrong for pairwise comparison.

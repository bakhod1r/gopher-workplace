# Escape Audit

## Intuition

Interfaces are not free. Storing a non-pointer value in an interface usually means copying it to the heap so the interface can hold a pointer to it — the cost shows up as allocations in a loop that looks identical to the concrete one.

## Approach

1. `SumValues` and `StackAgg.Sum` are plain loops over `[]int`; nothing escapes.
2. `SumBoxed` asserts each element with comma-ok and skips non-ints.
3. `BoxAll` converts explicitly, which is where the boxing allocations happen.
4. The tests measure the difference with `AllocsPerRun`.

## Solution

```go
func (StackAgg) Sum(vs []int) int {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	return sum
}

func SumValues(vs []int) int {
	sum := 0
	for _, v := range vs {
		sum += v
	}
	return sum
}

func SumBoxed(vs []any) int {
	sum := 0
	for _, v := range vs {
		if n, ok := v.(int); ok {
			sum += n
		}
	}
	return sum
}
```

## Walkthrough

`StackAgg` is an empty struct, so calling `Sum` through the `Summer` interface still allocates nothing — the interface itself is not the cost. `BoxAll` is: each `append(out, v)` boxes an int.

## Pitfalls

- Assuming any interface use allocates — the dispatch is cheap; the *boxing* of values is what costs.
- Writing a generic `Sum(vs []any)` API and paying conversion costs at every call site.
- Verifying escape claims by reading the code instead of measuring: use `go build -gcflags=-m` or `AllocsPerRun`.

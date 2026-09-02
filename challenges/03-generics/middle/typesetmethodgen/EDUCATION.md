# Type Set And Method

## Intuition

Combining a type set with a method is what lets one function do arithmetic and formatting without a type switch or reflection.

## Approach

1. Allocate the label slice.
2. For each element append `v.String()` and add `int(v)`.
3. Return both results.

## Solution

```go
func Labels[T Code](s []T) ([]string, int) {
	out := make([]string, 0, len(s))
	total := 0
	for _, v := range s {
		out = append(out, v.String())
		total += int(v)
	}
	return out, total
}
```

## Walkthrough

`Labels([]Status{200, 404})` calls each `String` and adds the underlying ints to 604.

## Pitfalls

- Dropping the `~int` half and losing the ability to convert.
- Dropping the method half and needing `fmt.Sprint` instead.
- Summing in `T` and returning it, which the signature does not ask for.

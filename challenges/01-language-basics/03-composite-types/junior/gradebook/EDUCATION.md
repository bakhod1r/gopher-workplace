# Maps of composite values

## Intuition

A map value can be any type, including a slice. `map[string][]int` associates
each key with a list:

```go
for name, scores := range book {
	if len(scores) == 0 { continue }
	sum := 0
	for _, s := range scores { sum += s }
	out[name] = sum / len(scores)
}
```

## Approach

1. make an empty result map.
2. Range students; skip any whose score slice is empty (avoids divide-by-zero and omits them).
3. Sum the slice, divide by its length with integer division, store under the name.
4. Return the map.

## Solution

```go
func Averages(book map[string][]int) map[string]int {
	result := make(map[string]int)
	for name, scores := range book {
		if len(scores) == 0 {
			continue
		}
		sum := 0
		for _, s := range scores {
			sum += s
		}
		result[name] = sum / len(scores)
	}
	return result
}
```

## Walkthrough

Averages: ann sums to 270, 270/3=90; bob sums to 145, 145/2=72; cid empty -> skipped. Result {"ann":90,"bob":72}.

## Pitfalls

- Skip empty slices before dividing.
- A missing map key of slice type reads as `nil` (len 0, safe to range).
- Integer average truncates.

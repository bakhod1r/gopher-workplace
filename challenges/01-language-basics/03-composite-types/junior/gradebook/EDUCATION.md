# Maps of composite values

## The idea

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

## Why it matters

Grouping — key to a collection — is fundamental (buckets, adjacency lists).
Guarding empty collections avoids divide-by-zero.

## Watch out

- Skip empty slices before dividing.
- A missing map key of slice type reads as `nil` (len 0, safe to range).
- Integer average truncates.

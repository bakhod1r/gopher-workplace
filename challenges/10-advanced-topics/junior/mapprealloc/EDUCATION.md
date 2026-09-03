# Size The Map Before Filling It

## Intuition

A map grows by allocating a bigger bucket array and migrating entries. If you know roughly how many keys are coming, saying so up front skips every intermediate size.

## Approach

1. `make(map[string]int, len(words))`.
2. Loop and `m[w]++`.
3. Return the map.

## Solution

```go
// Count returns how many times each word appears in words.
//
// The map must be created with a size hint so it does not rehash its way up
// from nothing while the loop runs.
//
// Examples:
//
// 	Count([]string{"a", "b", "a"}) => map[a:2 b:1]
func Count(words []string) map[string]int {
	m := make(map[string]int, len(words))
	for _, w := range words {
		m[w]++
	}
	return m
}
```

## Walkthrough

For 1000 tokens over 50 distinct words, an unsized map grows through several bucket-array sizes; a sized one starts big enough and never migrates.

## Pitfalls

- Over-hinting with a huge number wastes memory — `len(words)` is an upper bound, and that is fine here.
- The hint is advice, not a cap; the map still grows if you exceed it.

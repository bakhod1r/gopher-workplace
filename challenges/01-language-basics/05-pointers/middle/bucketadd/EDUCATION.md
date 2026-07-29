# Mutable map values via pointers

## Intuition

Map struct VALUES aren't addressable (`m[k].F = x` is illegal), so storing `*Struct` lets you fetch and mutate the bucket in place.

## Approach

1. Look up `m[key]`; a missing key yields a nil pointer.
2. Lazily create `&Bucket{}` on first use.
3. Add through the pointer: `m[key].Total += pts`.

## Solution

```go
type Bucket struct{ Total int }

func AddScore(m map[string]*Bucket, key string, pts int) {
	if m[key] == nil {
		m[key] = &Bucket{}
	}
	m[key].Total += pts
}
```

## Walkthrough

First `AddScore(m, "a", 3)` allocates a bucket with Total 3; the second call finds it and adds 4, reaching 7.

## Pitfalls

- `m[k].Total += x` on a map of struct VALUES won't compile.
- Store pointers and lazily create them.

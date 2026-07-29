# Pointers inside copied struct values

## Intuition

Copying a struct duplicates its pointer fields' addresses, so the copy still aliases the same pointee; mutate through the pointer even when the struct itself is a copy.

## Approach

1. The bug discards the record with `_ = r` and never mutates.
2. Follow the pointer field: `*r.P++`.

## Solution

```go
type Ref struct{ P *int }

func BumpVia(m map[int]Ref, k int) {
	r := m[k]
	*r.P++
}
```

## Walkthrough

`_ = r` reads the struct but changes nothing. `*r.P++` increments the int the record's pointer field references.

## Pitfalls

- `r := m[k]` copies the struct but not the pointee.
- `*r.P++` mutates the shared int.

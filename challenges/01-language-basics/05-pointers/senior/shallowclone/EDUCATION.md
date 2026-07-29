# Shallow vs deep struct copies

## Intuition

Copying a struct duplicates value fields but shares slices/maps/pointers; a deep clone must copy those reference fields explicitly.

## Approach

1. `cp := *d` copies the struct but the `Tags` slice header still shares the backing array.
2. Deep-copy the slice: `cp.Tags = append([]string(nil), d.Tags...)`.

## Solution

```go
type Doc struct {
	Tags []string
}

func Clone(d *Doc) *Doc {
	cp := *d
	cp.Tags = append([]string(nil), d.Tags...)
	return &cp
}
```

## Walkthrough

The shallow copy shares `Tags`, so mutating the clone's tags mutates the source. Cloning the slice into a fresh array breaks the aliasing.

## Pitfalls

- `*d` copies the header, not the slice's elements.
- Deep-copy each reference-typed field.

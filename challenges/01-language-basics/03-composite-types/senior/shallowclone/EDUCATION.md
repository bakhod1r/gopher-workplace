# Shallow vs deep copies

## Intuition

Copying a struct by value duplicates each field. For a slice field, that copies
the **header** (pointer/len/cap) but not the underlying array — so both structs
share the same elements:

```go
d.Tags = append([]string{}, d.Tags...) // independent copy of the slice
return d
```

## Approach

1. Bug: return d copies the struct but the Tags slice header still points at the same backing array, so the copy shares tags. 2. Fix: copy the struct, then replace Tags with a fresh slice: out.Tags = append([]string(nil), d.Tags...). 3. Now mutating one Doc's tags does not affect the other.

## Solution

```go
type Doc struct {
	Title string
	Tags  []string
}

func Clone(d Doc) Doc {
	out := d
	out.Tags = append([]string(nil), d.Tags...)
	return out
}
```

## Walkthrough

return d shares the Tags array: original.Tags[0]=z shows in the clone. Cloning Tags into a new array isolates the two.

## Pitfalls

- Every slice/map/pointer field needs its own deep copy.
- `slices.Clone`/`maps.Clone` handle one level; nested references need recursion.
- Value types (numbers, strings, arrays) copy fully and are safe.

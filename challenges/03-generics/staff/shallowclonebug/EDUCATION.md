# The Clone That Shares Its Tags

## Intuition

`out := d` duplicates the struct's bytes: `Title` becomes a genuine copy, but `Tags` is a pointer, length and capacity triple pointing at the original array. Both documents now write to the same tags.

## Approach

1. Copy the struct to get every scalar field.
2. Allocate a new slice for `Tags` at the same length.
3. Copy the elements into it.

## Solution

```go
func CloneDoc[T any](d Doc[T]) Doc[T] {
	out := d
	out.Tags = make([]T, len(d.Tags))
	copy(out.Tags, d.Tags)
	return out
}

func CloneAll[T any](ds []Doc[T]) []Doc[T] {
	out := make([]Doc[T], len(ds))
	for i, d := range ds {
		out[i] = CloneDoc(d)
	}
	return out
}
```

## Walkthrough

`c := CloneDoc(d); c.Tags[0] = "draft"` also sets `d.Tags[0]` to `"draft"`, because both headers address the same backing array.

## Pitfalls

- Using `out.Tags = d.Tags[:]`, which is the same aliasing with extra syntax.
- Forgetting that `append` to the clone may or may not alias, depending on spare capacity — a bug that appears intermittently.

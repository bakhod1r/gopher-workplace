# Mutating Slices with Value Receivers

## Intuition

A slice in Go is a header (pointer, length, capacity). When passed by value
`(l IntList)`, the header is copied, but the pointer still points to the same
backing array. Therefore, modifying `l[i]` modifies the original elements.

You only need a pointer receiver `*IntList` if you want to `append` and change
the length/capacity, or point to a completely different array.

## Approach

1. Range over indices and multiply.

## Solution

```go
func (l IntList) Apply(factor int) {
	for i := range l {
		l[i] *= factor
	}
}
```

## Walkthrough

For `IntList{1, 2, 3}`:
- `l[0] *= 2` → 2.
- `l[1] *= 2` → 4.
- Caller's slice header points to the same modified array.

## Pitfalls

- Using a pointer receiver `(l *IntList)` works but is less idiomatic unless
  appending.
- `for _, v := range l { v *= factor }` — `v` is a copy, so the slice elements
  are unchanged. Must use index `l[i]`.

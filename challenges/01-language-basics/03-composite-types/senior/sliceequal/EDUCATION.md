# Comparing slices

## Intuition

Slices are not comparable with `==` (except to `nil`). Equality is length plus
element-wise comparison:

```go
if len(a) != len(b) { return false }
for i := range a { if a[i] != b[i] { return false } }
return true
```

## Approach

1. Bug: after the length check the function just returns true, so any two equal-length slices compare equal regardless of contents. 2. Fix: loop over the elements and return false on the first mismatch, true if all match. 3. Length equality is necessary but not sufficient.

## Solution

```go
func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
```

## Walkthrough

a=[1,2,3], b=[1,9,3] have equal length; buggy code returns true. The element loop finds a[1]!=b[1] and returns false.

## Pitfalls

- `a == b` on slices is a compile error; don't reach for it.
- `slices.Equal` (Go 1.21+) does exactly this.
- `reflect.DeepEqual` also works but is slower and broader.

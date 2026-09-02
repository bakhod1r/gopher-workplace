# Interface Composition

## Intuition

Embedding interfaces builds bigger contracts out of small ones. Asserting to an interface asks a capability question at runtime: does this dynamic type have these methods?

## Approach

1. `File` gets pointer-receiver `Read` and `Write`.
2. In `Describe`, assert `v.(Reader)` and `v.(Writer)` with comma-ok, discarding the values.
3. Combine the two bools in a `switch` to pick the label.

## Solution

```go
func (f *File) Read() string { return f.data }

func (f *File) Write(s string) { f.data = s }

func Describe(v any) string {
	_, canRead := v.(Reader)
	_, canWrite := v.(Writer)
	switch {
	case canRead && canWrite:
		return "rw"
	case canRead:
		return "r"
	case canWrite:
		return "w"
	default:
		return "none"
	}
}
```

## Walkthrough

`Describe(File{})` — a value, not a pointer — returns `"none"`: pointer-receiver methods are not in the value's method set, so neither assertion succeeds.

## Pitfalls

- Asserting to `ReadWriter` only, which cannot distinguish `"r"` from `"w"`.
- Type-switching on the concrete types — `Describe` would then need editing for every new backend.
- Forgetting that a `File` value and a `*File` have different method sets.

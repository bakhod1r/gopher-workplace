# The Same Fields, Half The Bytes

## Intuition

The compiler will not reorder your fields, so the layout is your responsibility. Sorting from widest to narrowest lets each field land on its natural boundary with nothing skipped.

## Approach

1. Declare `Ref int64` first.
2. Then `Seq int32`.
3. Then the two `byte` fields.

## Solution

```go
import "unsafe"

// Entry is one cache record, ordered widest field first.
type Entry struct {
	Ref  int64
	Seq  int32
	Flag byte
	Kind byte
}

// Size returns the size of the Entry type.
//
// Entry's fields are declared in an order that forces the compiler to
// insert padding between them. Reordering them from widest to narrowest
// removes it without changing what the struct holds.
//
// Examples:
//
// 	Size() => 16 once the fields are ordered well
func Size() uintptr {
	return unsafe.Sizeof(Entry{})
}
```

## Walkthrough

As declared: Flag at 0, seven bytes of padding, Ref at 8, Kind at 16, three bytes of padding, Seq at 20 — 24 bytes. Reordered: Ref at 0, Seq at 8, Flag at 12, Kind at 13, then two bytes of tail padding — 16.

## Pitfalls

- Changing `int64` to `int32` to save space — that changes the type, and the test checks it.
- Assuming the saving is always this large; a struct of same-width fields has no padding to remove.

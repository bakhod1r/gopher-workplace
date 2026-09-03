# How Big Is This, Really

## Intuition

`Sizeof` answers "how many bytes does a variable of this type occupy", which is not the same as "how much memory does this value use". A string field is always two words, whatever text it points at.

## Approach

1. Declare a zero `Header`.
2. Return `unsafe.Sizeof` of the struct and of the two fields.

## Solution

```go
import "unsafe"

// Header is the fixed part of a record.
type Header struct {
	Id   int64
	Name string
	Tags []string
}

// Sizes returns the size in bytes of the Header type and of its Id and
// Name fields.
//
// unsafe.Sizeof is a compile-time constant: it measures the type, not the
// data a pointer or slice header refers to.
//
// Examples:
//
// 	Sizes() => 40, 8, 16 on a 64-bit build
func Sizes() (header, id, name uintptr) {
	var h Header
	return unsafe.Sizeof(h), unsafe.Sizeof(h.Id), unsafe.Sizeof(h.Name)
}
```

## Walkthrough

On a 64-bit build `Id` is 8 bytes, `Name` is 16 (pointer + length), `Tags` is 24 (pointer + length + capacity), so the struct is 48.

## Pitfalls

- Writing the numbers as literals — they change with the architecture.
- Expecting `Sizeof` on a slice to include the elements; it does not.

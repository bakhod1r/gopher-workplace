# Struct Layout

## Intuition

Every field must sit at an offset that is a multiple of its alignment. A `bool` before an `int64` forces seven bytes of padding; putting the `int64` first lets the two bools share the trailing word.

## Approach

1. Declare `Packed` with `B int64` first, then the two bools.
2. Both `Size` methods return `unsafe.Sizeof(p)` on the receiver.
3. `TotalBytes` multiplies the size by the count, clamping negatives.
4. Verify with `unsafe.Sizeof` and `unsafe.Offsetof` rather than by eye.

## Solution

```go
type Packed struct {
	B int64
	A bool
	C bool
}

func (p Padded) Size() uintptr { return unsafe.Sizeof(p) }

func (p Packed) Size() uintptr { return unsafe.Sizeof(p) }

func TotalBytes(s Sizer, n int) uintptr {
	if n < 0 {
		return 0
	}
	return s.Size() * uintptr(n)
}
```

## Walkthrough

`Padded` is 24 bytes: bool + 7 padding + int64 + bool + 7 padding. `Packed` is 16: int64 + bool + bool + 6 padding. Over 100M records that is 800MB saved by reordering three fields.

## Pitfalls

- Assuming the compiler reorders fields — Go never does; declaration order is the layout.
- Hardcoding sizes for one architecture: alignment differs on 32-bit builds.
- Reordering fields in a struct that is serialised positionally by `unsafe`, which changes the wire format.

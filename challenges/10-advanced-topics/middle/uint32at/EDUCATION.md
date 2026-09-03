# Read A Wide Value Out Of A Byte Buffer

## Intuition

Reinterpreting memory is only as safe as the checks you write around it. The compiler stops checking the moment you take an `unsafe.Pointer`, so both the range and the alignment become your responsibility.

## Approach

1. Reject `off < 0` and `off+4 > len(b)`.
2. `unsafe.Add` the data pointer by `off`.
3. Reject when the address's low two bits are set.
4. Read through `(*uint32)(p)`.

## Solution

```go
import "unsafe"

// Uint32At reads the native-endian uint32 at byte offset off in b.
//
// The read is bounds-checked and alignment-checked; anything out of range
// or misaligned reports false rather than faulting.
//
// Examples:
//
// 	Uint32At(buf, 0) => the first four bytes as a uint32
func Uint32At(b []byte, off int) (uint32, bool) {
	if off < 0 || off+4 > len(b) {
		return 0, false
	}
	p := unsafe.Add(unsafe.Pointer(unsafe.SliceData(b)), off)
	if uintptr(p)&3 != 0 {
		return 0, false
	}
	return *(*uint32)(p), true
}
```

## Walkthrough

For an 8-byte buffer and off = 4: the range check passes, the address is 4 bytes past an aligned base so the low bits are clear, and the read returns the second word.

## Pitfalls

- Checking `off < len(b)` instead of `off+4 <= len(b)` — that lets a read run three bytes past the end.
- Checking the alignment of `off` alone; the buffer's own start may be misaligned.

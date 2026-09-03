# Check The Alignment You Depend On

## Intuition

Hard-coded layout constants encode one machine's answer into code that will run on others. `Alignof` asks the compiler for the real requirement, so the check travels with the build.

## Approach

1. Take `want := unsafe.Alignof(c.Value)`.
2. Check that the field's offset is a multiple of `want`.
3. Check that an instance's field address is a multiple of `want`.

## Solution

```go
import "unsafe"

// Counter is a struct whose Value field is updated atomically.
type Counter struct {
	Value int64
	Name  string
}

// Check reports whether Counter's Value field is aligned well enough for
// 64-bit atomic operations.
//
// The requirement is the type's own alignment, which unsafe.Alignof
// reports. Hard-coding a number is how this check passes on the machine it
// was written on and nowhere else.
//
// Examples:
//
// 	Check() => true for a correctly laid out Counter
func Check() bool {
	var c Counter
	want := unsafe.Alignof(c.Value)
	return unsafe.Offsetof(c.Value)%want == 0 &&
		uintptr(unsafe.Pointer(&c.Value))%want == 0
}
```

## Walkthrough

On a 64-bit build `Alignof(int64)` is 8 and both checks pass. On a platform where it is 4, the literal 8 would have demanded more than the platform provides — the assertion would fail for the wrong reason.

## Pitfalls

- Checking only the offset; the struct itself may sit at a misaligned address inside another struct.
- Using `Sizeof` in place of `Alignof` — they agree for int64 and not in general.

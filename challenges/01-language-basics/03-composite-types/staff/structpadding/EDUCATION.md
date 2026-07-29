# Alignment and field ordering

## Intuition

Each field is aligned to its size; the compiler inserts padding to satisfy that,
and the struct's size rounds up to its largest alignment. `bool,int64,bool` pads
to **24** bytes; `int64,bool,bool` packs to **16**:

```go
type Record struct { B int64; A bool; C bool } // 8 + 1 + 1 -> padded to 16
```

## Approach

1. Bug: order A bool, B int64, C bool forces 7 bytes padding after A (to 8-align int64) plus tail padding -> 24 bytes.
2. Fix: put the 8-byte int64 first, then the two bools together: B int64, A bool, C bool -> 8 + 1 + 1 + 6 tail padding = 16 bytes.

## Solution

```go
type Record struct {
	B int64
	A bool
	C bool
}
```

## Walkthrough

On 64-bit, int64 needs 8-byte alignment. Bad layout: A(1)+pad(7)+B(8)+C(1)+pad(7)=24. Good layout: B(8)+A(1)+C(1)+pad(6)=16.

## Pitfalls

- Alignment is platform-dependent (this targets 64-bit).
- `unsafe.Sizeof` reveals the true size; `fieldalignment` (a vet tool) flags waste.
- Micro-optimize layout only for high-count structs; clarity first otherwise.

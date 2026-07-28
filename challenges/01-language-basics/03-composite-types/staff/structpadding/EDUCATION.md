# Alignment and field ordering

## The idea

Each field is aligned to its size; the compiler inserts padding to satisfy that,
and the struct's size rounds up to its largest alignment. `bool,int64,bool` pads
to **24** bytes; `int64,bool,bool` packs to **16**:

```go
type Record struct { B int64; A bool; C bool } // 8 + 1 + 1 -> padded to 16
```

## Why it matters

For large arrays of structs, field ordering directly controls memory and cache
footprint. Ordering fields from widest to narrowest minimizes padding — a real
performance lever in hot data structures.

## Watch out

- Alignment is platform-dependent (this targets 64-bit).
- `unsafe.Sizeof` reveals the true size; `fieldalignment` (a vet tool) flags waste.
- Micro-optimize layout only for high-count structs; clarity first otherwise.

# Reinterpreting memory with unsafe.Pointer

## The idea

`*(*T)(unsafe.Pointer(&x))` reads the same bytes as a different type; a value cast instead performs numeric conversion.

## Why it matters

Bit-level access (hashing, serialisation, `math.Float64bits`) reinterprets rather than converts.

## Watch out

- Narrowing to float32 changes the bit pattern and width.
- Reinterpret the float64 directly as uint64 (both 8 bytes).

# Width matching in reinterprets

## The idea

A reinterpret cast must use a target type the same size as the source; a narrower type reads only part of the memory.

## Why it matters

Mismatched-width reinterprets silently truncate or read garbage.

## Watch out

- int64 is 8 bytes; `*uint32` reads only the low 4.
- Match widths: `*uint64` for an int64.

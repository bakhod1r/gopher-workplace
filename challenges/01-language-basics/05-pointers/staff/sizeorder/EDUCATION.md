# Struct layout and padding

## The idea

The compiler pads fields to satisfy alignment and rounds the struct to its largest alignment; ordering fields widest-first minimises wasted padding.

## Why it matters

For large arrays of structs, layout directly controls memory and cache footprint.

## Watch out

- `bool,int64,bool` -> 24 bytes; `int64,bool,bool` -> 16.
- Alignment is platform-dependent; this targets 64-bit.

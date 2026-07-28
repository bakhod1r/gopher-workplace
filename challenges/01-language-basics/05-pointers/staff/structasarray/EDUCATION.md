# Reinterpreting structs as arrays

## The idea

A struct of homogeneous fields has the same memory layout as an array of that type, so an unsafe reinterpret gives an indexable view of all fields.

## Why it matters

Bulk/SIMD-style field access reinterprets structs as arrays when layout matches.

## Watch out

- The reinterpreted array has one element per field.
- Reading only arr[0] ignores the rest of the struct.

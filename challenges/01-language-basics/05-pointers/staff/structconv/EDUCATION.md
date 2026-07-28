# Reinterpreting compatible structs

## The idea

Structs with identical memory layout can be converted with a single unsafe reinterpret, copying all fields at once.

## Why it matters

Layout-compatible reinterprets avoid field-by-field copies in interop code.

## Watch out

- `*(*Vec)(unsafe.Pointer(p))` copies X and Y together.
- Rebuilding by hand risks dropping fields.

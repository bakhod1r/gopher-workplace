# Slice header vs backing data

## The idea

A slice value is a header; `&s` addresses that header, while `&s[0]` (or `unsafe.SliceData`) addresses the actual data.

## Why it matters

Interop and zero-copy code must target the data pointer, not the header.

## Watch out

- `&s` writes over the slice header fields, not the elements.
- `&s[0]` is the data pointer.

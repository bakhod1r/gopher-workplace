# String header vs data pointer

## The idea

A string is a (data, len) header; `&s` addresses that header, while `unsafe.StringData(s)` addresses the bytes.

## Why it matters

Zero-copy string/byte interop must target the data pointer.

## Watch out

- `&s` reads the header fields, not the characters.
- `unsafe.StringData(s)` is the data pointer (bytes are immutable).

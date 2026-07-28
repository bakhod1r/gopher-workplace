# Sizing pointer-holding structs

## The idea

Each pointer field is word-sized; the struct's size is the sum (plus any padding). Measure the struct type, not a single pointer.

## Why it matters

Correct struct sizing matters for allocation and cache reasoning.

## Watch out

- One pointer is 8 bytes; two make 16.
- `Sizeof(Pair{})` measures the whole struct.

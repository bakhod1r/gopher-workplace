# Distinct allocations per iteration

## The idea

Taking the address of a single hoisted variable shares one allocation; declaring the variable inside the loop gives each pointer its own heap object.

## Why it matters

Reusing one struct's address across iterations is a real aliasing bug in builders.

## Watch out

- `&it` of a hoisted `it` aliases one struct.
- Declare `it := Item{...}` inside the loop for distinct pointers.

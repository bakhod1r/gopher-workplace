# Mutating pointees fetched from maps

## The idea

Dereferencing into a local copies the struct; to change the stored object, write through the fetched pointer.

## Why it matters

Copy-then-mutate loses updates to pointer-valued map entries.

## Watch out

- `acc := *a; acc.Balance += amt` edits a throwaway copy.
- `a.Balance += amt` edits the stored account.

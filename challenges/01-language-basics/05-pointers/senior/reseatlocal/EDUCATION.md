# Double-pointer reassignment

## The idea

`pp` is a copy of the address of the caller's pointer; only `*pp = q` writes through to reseat it.

## Why it matters

Reassigning the parameter instead of `*pp` is a silent no-op in tree/list update APIs.

## Watch out

- `pp = &q` changes a local copy.
- `*pp = q` reseats the caller's pointer.

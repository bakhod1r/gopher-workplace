# Element pointers across reallocation

## The idea

`append` beyond capacity detaches pre-existing element pointers from the live slice; you must re-take the address after growth.

## Why it matters

Holding element pointers across appends corrupts buffer/pool code that grows in place.

## Watch out

- After a reallocating append, `&xs[i]` from before is stale.
- Re-take it, or index the current slice.

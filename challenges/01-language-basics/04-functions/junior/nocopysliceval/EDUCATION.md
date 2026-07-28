# Slices are references under a value header

## The idea

Passing a slice copies its (ptr,len,cap) header but not the backing array, so element writes are visible to the caller — reads are always safe.

## Why it matters

Knowing what is copied vs shared prevents accidental mutation of a caller's data.

## Watch out

- Don't sort or assign into `xs[i]` — that would mutate the caller's array.
- The named return `sum` starts at 0.

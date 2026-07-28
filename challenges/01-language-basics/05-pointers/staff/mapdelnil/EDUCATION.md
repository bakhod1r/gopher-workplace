# Deleting map entries

## The idea

Assigning a nil (or zero) value keeps the key; the `delete` built-in is the only way to remove an entry from a map.

## Why it matters

Nil-ing instead of deleting leaves phantom keys that skew counts and iteration.

## Watch out

- `m[id] = nil` leaves the key present with a nil value.
- `delete(m, id)` removes it entirely.

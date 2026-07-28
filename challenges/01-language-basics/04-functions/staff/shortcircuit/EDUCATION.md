# Short-circuit evaluation as a guard

## The idea

`&&` and `||` evaluate left to right and stop early; the nil/bounds guard must precede the operation it protects.

## Why it matters

Reordering a short-circuit guard reintroduces the nil-deref (or index) panic it was meant to prevent.

## Watch out

- `p != nil && *p > 0` is safe; the reverse panics on nil.
- Rely on left-to-right short-circuit ordering deliberately.

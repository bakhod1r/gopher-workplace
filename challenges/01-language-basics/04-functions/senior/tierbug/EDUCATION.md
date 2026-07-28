# Explicit fallthrough in Go

## The idea

Unlike C, Go breaks after each case; `fallthrough` opts in and ignores the next case's condition — misuse silently misroutes values.

## Why it matters

Porting C/Java switch code often carries over unwanted fallthrough assumptions.

## Watch out

- Only add `fallthrough` when you truly want the next case body.
- Code after an unconditional `return` is dead.

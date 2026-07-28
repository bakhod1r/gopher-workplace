# Value vs pointer receivers

## The idea

A value receiver copies the struct, so any field write inside the method affects
only the copy:

```go
func (w *Wallet) Credit(amount int) { w.Balance += amount } // pointer: mutates
```

## Why it matters

This compiles cleanly and looks correct, but state never updates — a classic Go
bug. Any method that mutates must use a pointer receiver.

## Watch out

- `w.Credit(...)` auto-takes the address when `w` is addressable, hiding the
  distinction.
- Keep a type's receivers consistent (all pointer if any mutate).
- Value receivers are right for small, immutable value types.

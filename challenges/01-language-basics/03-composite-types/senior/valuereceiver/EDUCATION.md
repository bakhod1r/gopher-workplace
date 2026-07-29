# Value vs pointer receivers

## Intuition

A value receiver copies the struct, so any field write inside the method affects
only the copy:

```go
func (w *Wallet) Credit(amount int) { w.Balance += amount } // pointer: mutates
```

## Approach

1. Bug: func (w Wallet) Credit uses a value receiver, so w is a copy; w.Balance += amount mutates the copy and is lost. 2. Fix: use a pointer receiver func (w *Wallet) Credit. 3. Then the method mutates the caller's Wallet in place.

## Solution

```go
type Wallet struct {
	Balance int
}

func (w *Wallet) Credit(amount int) {
	w.Balance += amount
}
```

## Walkthrough

With a value receiver, Credit(100) updates a copy; the caller's Balance stays 0. A pointer receiver writes through to the original -> Balance==100.

## Pitfalls

- `w.Credit(...)` auto-takes the address when `w` is addressable, hiding the
  distinction.
- Keep a type's receivers consistent (all pointer if any mutate).
- Value receivers are right for small, immutable value types.

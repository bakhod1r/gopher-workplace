# Pointer receivers mutate

## Intuition

A method with a **pointer receiver** operates on the original struct:

```go
func (a *Account) Deposit(amount int) { a.Balance += amount }
```

A value receiver would receive a copy and its changes would vanish.

## Approach

1. Use the pointer receiver (a *Account) so the method mutates the caller's struct.
2. a.Balance += amount adds to the stored balance in place.

## Solution

```go
type Account struct {
	Balance int
}

func (a *Account) Deposit(amount int) {
	a.Balance += amount
}
```

## Walkthrough

a=&Account{100}: Deposit(50) -> 150; Deposit(25) -> 175. Same underlying struct is updated each time.

## Pitfalls

- Pointer-receiver calls need an **addressable** value; a map element is not.
- Value receiver + field write silently mutates a copy.
- Large structs are cheaper to pass by pointer regardless of mutation.
